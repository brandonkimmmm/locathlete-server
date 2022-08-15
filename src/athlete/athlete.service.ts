import { Injectable, Logger } from '@nestjs/common';
import { Athlete, AthletesOnSchools, School, SchoolType } from '@prisma/client';
import {
	getOrderingQuery,
	getPaginationQuery
} from 'src/common/helpers/query.helper';
import { PrismaService } from 'src/prisma/prisma.service';
import { GetAthletesDTO } from './dto';

@Injectable()
export class AthleteService {
	private readonly logger: Logger = new Logger(AthleteService.name);

	constructor(private readonly prismaService: PrismaService) {}

	async findAllAthletesPaginated(dto: GetAthletesDTO) {
		this.logger.debug('findAllAthletesPaginated dto', dto);
		const { limit, page, order_by, order } = dto;
		const count = await this.prismaService.athlete.count();
		const rows = await this.prismaService.athlete.findMany({
			...getPaginationQuery(limit, page),
			...getOrderingQuery(order_by, order),
			include: {
				schools: {
					include: {
						school: {
							include: {
								school_type: true
							}
						}
					}
				}
			}
		});
		this.logger.debug('findAllAthletesPaginated count', { count });
		return {
			count,
			rows
		};
	}

	serializeAthlete(
		data: Athlete & {
			schools: (AthletesOnSchools & {
				school: School & {
					school_type: SchoolType;
				};
			})[];
		}
	) {
		const { id, bio, first_name, middle_name, last_name } = data;
		return {
			id,
			bio,
			first_name,
			middle_name,
			last_name,
			full_name: `${first_name} ${
				middle_name ? `${middle_name} ${last_name}` : last_name
			}`
		};
	}
}
