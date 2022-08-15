import { Controller, Get, Logger, Query } from '@nestjs/common';
import { AthleteService } from './athlete.service';
import { GetAthletesDTO } from './dto';

@Controller('athletes')
export class AthleteController {
	private readonly logger: Logger = new Logger(AthleteController.name);

	constructor(private readonly athleteService: AthleteService) {}

	@Get()
	async getAthletes(@Query() dto: GetAthletesDTO) {
		this.logger.debug('getAthletes request dto', dto);
		const { count, rows } =
			await this.athleteService.findAllAthletesPaginated(dto);
		return {
			count,
			rows: rows.map(this.athleteService.serializeAthlete)
		};
	}
}
