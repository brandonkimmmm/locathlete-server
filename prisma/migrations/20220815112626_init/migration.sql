-- CreateTable
CREATE TABLE "athletes" (
    "id" SERIAL NOT NULL,
    "bio" TEXT NOT NULL,
    "first_name" VARCHAR(255) NOT NULL,
    "middle_name" VARCHAR(255),
    "last_name" VARCHAR(255) NOT NULL,
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT "athletes_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "schools" (
    "id" SERIAL NOT NULL,
    "description" TEXT NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "street_address" VARCHAR(255) NOT NULL,
    "city" VARCHAR(255) NOT NULL,
    "country" VARCHAR(255) NOT NULL,
    "administration_area" VARCHAR(255),
    "postal_code" VARCHAR(255) NOT NULL,
    "lat" DOUBLE PRECISION NOT NULL,
    "lng" DOUBLE PRECISION NOT NULL,
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "school_type_id" INTEGER NOT NULL,

    CONSTRAINT "schools_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "school_types" (
    "id" SERIAL NOT NULL,
    "title" VARCHAR(255) NOT NULL,
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT "school_types_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "athletes_on_schools" (
    "athlete_id" INTEGER NOT NULL,
    "school_id" INTEGER NOT NULL,
    "start_date" DATE NOT NULL,
    "end_date" DATE NOT NULL,

    CONSTRAINT "athletes_on_schools_pkey" PRIMARY KEY ("athlete_id","school_id")
);

-- CreateIndex
CREATE UNIQUE INDEX "school_types_title_key" ON "school_types"("title");

-- AddForeignKey
ALTER TABLE "schools" ADD CONSTRAINT "schools_school_type_id_fkey" FOREIGN KEY ("school_type_id") REFERENCES "school_types"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "athletes_on_schools" ADD CONSTRAINT "athletes_on_schools_athlete_id_fkey" FOREIGN KEY ("athlete_id") REFERENCES "athletes"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "athletes_on_schools" ADD CONSTRAINT "athletes_on_schools_school_id_fkey" FOREIGN KEY ("school_id") REFERENCES "schools"("id") ON DELETE CASCADE ON UPDATE CASCADE;
