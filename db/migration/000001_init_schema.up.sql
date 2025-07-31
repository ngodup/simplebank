CREATE TABLE IF NOT EXISTS "accounts" (
  "id" bigserial PRIMARY KEY,
  "owner" varchar NOT NULL,
  "balance" bigint NOT NULL,
  "currency" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS "entries" (
  "id" bigserial PRIMARY KEY,
  "account_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS "transfers" (
  "id" bigserial PRIMARY KEY,
  "from_account_id" bigint NOT NULL,
  "to_account_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now()
);

-- ✅ Correct CREATE INDEX statements
CREATE INDEX IF NOT EXISTS idx_accounts_owner ON "accounts" ("owner");
CREATE INDEX IF NOT EXISTS idx_entries_account_id ON "entries" ("account_id");
CREATE INDEX IF NOT EXISTS idx_transfers_from_account_id ON "transfers" ("from_account_id");
CREATE INDEX IF NOT EXISTS idx_transfers_to_account_id ON "transfers" ("to_account_id");
CREATE INDEX IF NOT EXISTS idx_transfers_from_to ON "transfers" ("from_account_id", "to_account_id");

-- ✅ Comments
COMMENT ON COLUMN "entries"."amount" IS 'can be negative or positive';
COMMENT ON COLUMN "transfers"."amount" IS 'must be positive';

-- ✅ Foreign Keys
ALTER TABLE "entries"
  ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers"
  ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers"
  ADD FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("id");
