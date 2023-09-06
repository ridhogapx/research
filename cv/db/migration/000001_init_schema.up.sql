CREATE TABLE "cvzn_users" (
  "email" VARCHAR PRIMARY KEY NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "cvzn_personal_data" (
  "id" VARCHAR PRIMARY KEY NOT NULL,
  "email_user" VARCHAR NOT NULL,
  "full_name" VARCHAR,
  "mobile_phone_number" VARCHAR,
  "date_of_birth" VARCHAR,
  "image_url" VARCHAR,
  "address" VARCHAR,
  "website" VARCHAR,
  "linkedln" VARCHAR,
  "instagram" VARCHAR,
  "facebook" VARCHAR,
  "about_me" VARCHAR
);

CREATE TABLE "cvzn_history_education" (
  "id" VARCHAR PRIMARY KEY NOT NULL,
  "email_user" VARCHAR NOT NULL,
  "agency_name" VARCHAR,
  "major" VARCHAR,
  "start_date" VARCHAR,
  "date_of_completion" VARCHAR,
  "description" VARCHAR
);

CREATE TABLE "cvzn_work_experience" (
  "id" VARCHAR PRIMARY KEY NOT NULL,
  "email_user" VARCHAR NOT NULL,
  "position" VARCHAR,
  "company_name" VARCHAR,
  "address" VARCHAR,
  "start_date" VARCHAR,
  "date_of_completion" VARCHAR,
  "job_type" VARCHAR,
  "description" VARCHAR
);

CREATE TABLE "cvzn_expertise" (
  "id" VARCHAR PRIMARY KEY NOT NULL,
  "user_email" VARCHAR NOT NULL,
  "skill_name" VARCHAR
);

CREATE TABLE "cvzn_certification" (
  "id" VARCHAR PRIMARY KEY NOT NULL,
  "user_email" VARCHAR NOT NULL,
  "certificate_name" VARCHAR,
  "publisher" VARCHAR,
  "issued" VARCHAR,
  "certification_description" VARCHAR
);

CREATE TABLE "cvzn_organization" (
  "id" VARCHAR PRIMARY KEY NOT NULL,
  "user_email" VARCHAR NOT NULL,
  "organization_name" VARCHAR,
  "position_name" VARCHAR,
  "start_date" VARCHAR,
  "date_of_completion" VARCHAR,
  "organization_description" VARCHAR
);

CREATE TABLE "cvzn_hobby" (
  "id" VARCHAR PRIMARY KEY NOT NULL,
  "cvz_personal_data_id" VARCHAR NOT NULL,
  "hobby_name" VARCHAR
);

ALTER TABLE "cvzn_personal_data" ADD FOREIGN KEY ("email_user") REFERENCES "cvzn_users" ("email");

ALTER TABLE "cvzn_history_education" ADD FOREIGN KEY ("email_user") REFERENCES "cvzn_users" ("email");

ALTER TABLE "cvzn_work_experience" ADD FOREIGN KEY ("email_user") REFERENCES "cvzn_users" ("email");

ALTER TABLE "cvzn_expertise" ADD FOREIGN KEY ("user_email") REFERENCES "cvzn_users" ("email");

ALTER TABLE "cvzn_certification" ADD FOREIGN KEY ("user_email") REFERENCES "cvzn_users" ("email");

ALTER TABLE "cvzn_organization" ADD FOREIGN KEY ("user_email") REFERENCES "cvzn_users" ("email");

ALTER TABLE "cvzn_hobby" ADD FOREIGN KEY ("cvz_personal_data_id") REFERENCES "cvzn_users" ("email");
