CREATE TABLE IF NOT EXISTS products(
    "id" SERIAL PRIMARY KEY,
    "name" VARCHAR(255),
    "sku" VARCHAR(255),
    "description" VARCHAR(255),
    "price" DECIMAL(18,2),
    "count" INTEGER,
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
