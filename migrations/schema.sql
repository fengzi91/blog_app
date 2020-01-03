CREATE TABLE IF NOT EXISTS "schema_migration" (
"version" TEXT NOT NULL
);
CREATE UNIQUE INDEX "schema_migration_version_idx" ON "schema_migration" (version);
CREATE TABLE IF NOT EXISTS "users" (
"id" TEXT PRIMARY KEY,
"username" TEXT NOT NULL,
"email" TEXT NOT NULL,
"admin" NUMERIC NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
, "password_hash" TEXT NOT NULL DEFAULT '');
CREATE TABLE IF NOT EXISTS "posts" (
"id" TEXT PRIMARY KEY,
"title" TEXT NOT NULL,
"content" TEXT NOT NULL,
"author_id" char(36) NOT NULL,
"category_id" char(36) NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
, "attachment_id" char(36) NOT NULL DEFAULT '0393a648-5a61-4954-8fe6-a13e8de97b23', "subject" TEXT NOT NULL DEFAULT '');
CREATE TABLE IF NOT EXISTS "comments" (
"id" TEXT PRIMARY KEY,
"content" TEXT NOT NULL,
"author_id" char(36) NOT NULL,
"post_id" char(36) NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
CREATE TABLE IF NOT EXISTS "categories" (
"id" TEXT PRIMARY KEY,
"name" TEXT NOT NULL,
"slug" TEXT NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
, "orders" INTEGER NOT NULL DEFAULT '0');
CREATE TABLE IF NOT EXISTS "attachments" (
"id" TEXT PRIMARY KEY,
"author_id" char(36) NOT NULL,
"url" TEXT NOT NULL,
"size" INTEGER NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
