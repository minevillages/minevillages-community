
CREATE TABLE "post"(
    "postId" VARCHAR(255) NOT NULL,
    "pageId" VARCHAR(255) NOT NULL,
    "serverUserId" VARCHAR(255) NOT NULL,
    "content" VARCHAR(255) NOT NULL,
    "image" VARCHAR(255) NOT NULL,
    "createdAt" DATE NOT NULL,
    "updatedAt" DATE NOT NULL,
    "deletedAt" DATE NOT NULL,
    PRIMARY KEY("postId")
);

CREATE TABLE "emojji"(
    "emojiId" VARCHAR(255) NOT NULL,
    "postId" VARCHAR(255) NOT NULL,
    "serverUserId" VARCHAR(255) NOT NULL,
    "content" VARCHAR(255) NOT NULL,
    PRIMARY KEY("emojiId")
);
CREATE TABLE "page"(
    "pageId" VARCHAR(255) NOT NULL,
    "title" VARCHAR(255) NOT NULL,
    "permission" VARCHAR(255) NOT NULL,
    "createdAt" DATE NOT NULL,
    "updatedAt" DATE NOT NULL,
    "deletedAt" DATE NOT NULL,
    PRIMARY KEY("pageId")
);
ALTER TABLE
    "emojji" ADD CONSTRAINT "emojji_postid_foreign" FOREIGN KEY("postId") REFERENCES "post"("postId");
ALTER TABLE
    "post" ADD CONSTRAINT "post_pageid_foreign" FOREIGN KEY("pageId") REFERENCES "page"("pageId");