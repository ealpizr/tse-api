package handlers

import (
	"context"
	"log"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"github.com/ealpizr/tse-api/src/consts"
	"github.com/ealpizr/tse-api/src/models"
	"github.com/ealpizr/tse-api/src/utils"
	"github.com/gofiber/fiber/v2"
)

// @Summary      Query by full name
// @Description  get string by full name
// @Accept       json
// @Produce      json
// @Param        fullName   path      string  true  "Full name"
// @Success      200  {object}  models.Record
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /name [post]

func FindByFullNameRequestHandler(c *fiber.Ctx) error {
	r := new(models.Record)
	c.BodyParser(r)

	if r.FullName == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "fullName is required"})
	}

	if !utils.IsFullNameValid(r.FullName) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "given fullName is invalid"})
	}

	firstName, fLastName, sLastName := utils.ParseFullName(r.FullName)

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	err := chromedp.Run(ctx,
		chromedp.Navigate(consts.TSE_NAME_QUERY_URL),
		chromedp.WaitVisible(consts.HTML_INPUT_NAME, chromedp.ByID),
		chromedp.SendKeys(consts.HTML_INPUT_NAME, firstName, chromedp.ByID),
		chromedp.SendKeys(consts.HTML_INPUT_FLASTNAME, fLastName, chromedp.ByID),
		chromedp.SendKeys(consts.HTML_INPUT_SLASTNAME, sLastName, chromedp.ByID),
	)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "something wrong happened"})
	}

	// Actions like click don't always trigger a page navigation, so they don't wait
	// for a page load. Wrapping them with RunResponse does that waiting.
	_, err = chromedp.RunResponse(ctx, chromedp.Click(consts.HTML_BUTTON_NAME_QUERY, chromedp.ByID))
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "something wrong happened"})
	}

	// Check for HTML_LABEL_NOT_FOUND
	nodes := []*cdp.Node{}
	err = chromedp.Run(ctx, chromedp.Nodes(consts.HTML_LABEL_NOT_FOUND, &nodes, chromedp.AtLeast(0)))
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "something wrong happened"})
	}
	if len(nodes) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "no results found"})
	}

	err = chromedp.Run(ctx,
		chromedp.Nodes(consts.HTML_LABEL_RESULT_ENTRY, &nodes, chromedp.ByQueryAll),
	)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "something wrong happened"})
	}

	return c.Status(fiber.StatusOK).JSON(utils.ParseFoundResults(nodes))
}
