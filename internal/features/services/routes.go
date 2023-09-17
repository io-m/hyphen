package services

/*
GET /services - List All Services
POST /services - Create a Service
PUT /services/{serviceId} - Edit a Service
DELETE /services/{serviceId} - Delete a Service
GET /services/{serviceId}/availability - Search by Service with query parameters (e.g., ?date=2023-10-10)
GET /services/{serviceId} - Search by Service with query parameters (e.g., ?limit=10&latitude=45,123453&longitude=23,214533&radius=30)

SERVICE TAGS & CATEGORIES
POST /services/categories - Add a Category
PUT /services/categories/{categoryId} - Edit a Category
DELETE /services/categories/{categoryId} - Delete a Category
GET /services/categories - List All Categories
POST /services/tags - Add a Tag
PUT /services/tags/{tagId} - Edit a Tag
DELETE /services/tags/{tagId} - Delete a Tag
GET /services/tags - List All Tags

IMAGE repo dependency
POST /services/{serviceId}/images - Add Images to Service
DELETE /services/{serviceId}/images/{imageId} - Delete a Service Image
*/
