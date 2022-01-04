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

func FindByIDRequestHandler(c *fiber.Ctx) error {
	r := new(models.Record)
	c.BodyParser(r)

	if r.ID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "id is required"})
	}

	if !utils.IsIDValid(r.ID) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "given id is invalid"})
	}

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	err := chromedp.Run(ctx,
		chromedp.Navigate(consts.TSE_ID_QUERY_URL),
		chromedp.SendKeys(consts.HTML_INPUT_ID, r.ID, chromedp.ByID),
	)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "something wrong happened"})
	}

	// Actions like click don't always trigger a page navigation, so they don't wait
	// for a page load. Wrapping them with RunResponse does that waiting.
	_, err = chromedp.RunResponse(ctx, chromedp.Click(consts.HTML_BUTTON_ID_QUERY, chromedp.ByID))
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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "no results were found"})
	}

	err = chromedp.Run(ctx,
		chromedp.Text(consts.HTML_LABEL_FULLNAME, &r.FullName, chromedp.ByID),
		chromedp.Text(consts.HTML_LABEL_BIRTHDATE, &r.BirthDate, chromedp.ByID),
		chromedp.Text(consts.HTML_LABEL_AGE, &r.Age, chromedp.ByID),
	)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "something wrong happened"})
	}

	return c.Status(fiber.StatusOK).JSON(r)
}
