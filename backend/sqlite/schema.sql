    CREATE TABLE IF NOT EXISTS "users" (
        "user_id" INTEGER PRIMARY KEY AUTOINCREMENT,
        "nick_name" TEXT NOT NULL UNIQUE,
        "age" INTEGER NOT NULL CHECK (age >= 18),
        "gender" TEXT NOT NULL,
        "first_name" TEXT NOT NULL,
        "last_name" TEXT NOT NULL,
        "email" TEXT NOT NULL UNIQUE,
        "password_hash" TEXT NOT NULL,
        "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

    CREATE TABLE IF NOT EXISTS "posts" (
        "post_id" INTEGER PRIMARY KEY AUTOINCREMENT,
        "user_id" INTEGER NOT NULL,
        "title" TEXT NOT NULL,
        "content" TEXT NOT NULL,
        "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE
    );
    
    CREATE TABLE IF NOT EXISTS "comments" (
        "comments_id" INTEGER PRIMARY KEY AUTOINCREMENT,
        "post_id" INTEGER NOT NULL,
        "user_id" INTEGER NOT NULL,
        "content" TEXT NOT NULL,
        "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY ("post_id") REFERENCES "posts" ("id") ON DELETE CASCADE,
        FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE
    );

    CREATE TABLE IF NOT EXISTS "likes" (
        "likes_id" INTEGER PRIMARY KEY AUTOINCREMENT,
        "post_id" INTEGER NOT NULL,
        "user_id" INTEGER NOT NULL,
        "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY ("post_id") REFERENCES "posts" ("id") ON DELETE CASCADE,
        FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE
    );

    CREATE TABLE IF NOT EXISTS "dislikes" (
    "dislikes_id" INTEGER PRIMARY KEY AUTOINCREMENT,
    FOREIGN KEY ("post_id") REFERENCES "posts" ("id") ON DELETE CASCADE,
    FOREIGN KEY ("comment_id") REFERENCES "comments" ("id") ON DELETE CASCADE,
    FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE
    );

    CREATE TABLE IF NOT EXISTS "messages" (
        "messages_id" INTEGER PRIMARY KEY AUTOINCREMENT,
        "sender_id" INTEGER NOT NULL,
        "receiver_id" INTEGER NOT NULL,
        "content" TEXT NOT NULL,
        "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY ("sender_id") REFERENCES "users" ("id") ON DELETE CASCADE,
        FOREIGN KEY ("receiver_id") REFERENCES "users" ("id") ON DELETE CASCADE
    );