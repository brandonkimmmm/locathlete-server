import dotenv from 'dotenv';
import { PrismaClient } from '@prisma/client';
import athletes from './athletes';
import schoolTypes from './schoolTypes';

const prisma = new PrismaClient();

dotenv.config();

(async () => {
	try {
		console.log('Creating School Types...');
		for (const type of schoolTypes) {
			await prisma.schoolType.upsert({
				where: {
					title: type
				},
				update: {},
				create: {
					title: type
				}
			});
		}
		console.log('School Types Created!');

		console.log('Adding athletes and schools...');
		for (const data of athletes) {
			const { schools, ...athlete } = data;
			let dbAthlete = await prisma.athlete.findFirst({
				where: {
					first_name: athlete.first_name,
					last_name: athlete.last_name,
					middle_name: athlete.middle_name
				}
			});
			if (!dbAthlete) {
				dbAthlete = await prisma.athlete.create({
					data: {
						...athlete
					}
				});
			}
			for (const school of schools) {
				const { start_date, end_date, type, ...data } = school;
				let dbSchool = await prisma.school.findFirst({
					where: {
						name: data.name
					}
				});
				if (!dbSchool) {
					dbSchool = await prisma.school.create({
						data: {
							...data,
							school_type: {
								connect: {
									title: type
								}
							}
						}
					});
				}
				const athleteSchool = await prisma.athletesOnSchools.findFirst({
					where: {
						athlete_id: dbAthlete.id,
						school_id: dbSchool.id
					}
				});

				if (!athleteSchool) {
					await prisma.athletesOnSchools.create({
						data: {
							athlete: {
								connect: { id: dbAthlete.id }
							},
							school: {
								connect: { id: dbSchool.id }
							},
							start_date,
							end_date
						}
					});
				}
			}
		}
		console.log('Athletes and schools created!');
	} catch (err) {
		console.error(err);
		process.exit(1);
	} finally {
		await prisma.$disconnect();
		process.exit(0);
	}
})();
