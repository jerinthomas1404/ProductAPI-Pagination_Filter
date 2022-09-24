package controllers

import (
	"context"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/gofiber/fiber/v2"
	"github.com/jeirnthomas1404/ProductAPI-Pagination_Filter/configs"
	"github.com/jeirnthomas1404/ProductAPI-Pagination_Filter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetFilteredData(client *mongo.Client) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		collection := configs.GetCollection(client, "products")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var products []models.Product

		filter := bson.M{}
		if s := c.Query("s"); s != "" {
			filter = bson.M{
				"$or": []bson.M{
					{
						"title": bson.M{
							"$regex": primitive.Regex{
								Pattern: s,
								Options: "i",
							},
						},
					},
					{
						"description": bson.M{
							"$regex": primitive.Regex{
								Pattern: s,
								Options: "i",
							},
						},
					},
				},
			}
		}

		findOptions := options.Find()

		if sort := c.Query("sort"); sort != "" {
			if sort == "asc" {
				findOptions.SetSort(bson.D{{Key: "price", Value: 1}})
			} else if sort == "desc" {
				findOptions.SetSort(bson.D{{Key: "price", Value: -1}})
			}
		}

		page, _ := strconv.Atoi(c.Query("page", "1"))
		var perPage int64 = 10

		total, _ := collection.CountDocuments(ctx, filter)
		findOptions.SetSkip((int64(page) - 1) * perPage)
		findOptions.SetLimit(perPage)

		cursor, _ := collection.Find(ctx, filter, findOptions)
		defer cursor.Close(ctx)

		for cursor.Next(ctx) {
			var prd models.Product
			cursor.Decode(&prd)
			products = append(products, prd)
		}
		return c.JSON(fiber.Map{
			"data":      products,
			"total":     total,
			"page":      page,
			"last_page": math.Ceil(float64(total / perPage)),
		})

	}
}

func GetData(client *mongo.Client) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		collection := configs.GetCollection(client, "products")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var products []models.Product

		cursor, _ := collection.Find(ctx, bson.M{})
		defer cursor.Close(ctx)

		for cursor.Next(ctx) {
			var prd models.Product
			cursor.Decode(&prd)
			products = append(products, prd)
		}
		return c.JSON(products)

	}
}

func PopulateData(client *mongo.Client) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		collection := configs.GetCollection(client, "products")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		for i := 0; i < 50; i++ {
			var prd = &models.Product{
				Id:          primitive.NewObjectID(),
				Title:       faker.Word(),
				Description: faker.Paragraph(),
				Image:       fmt.Sprintf("http://lorempixel.com/200/200/?%s", faker.UUIDDigit()),
				Price:       rand.Intn(90) + 110,
			}
			collection.InsertOne(ctx, prd)

		}
		return c.JSON(fiber.Map{
			"message": "success",
		})
	}
}
