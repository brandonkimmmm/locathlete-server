import {
	INestApplication,
	Injectable,
	Logger,
	OnModuleInit
} from '@nestjs/common';
import { PrismaClient } from '@prisma/client';
import { ConfigService } from '@nestjs/config';

@Injectable()
export class PrismaService extends PrismaClient implements OnModuleInit {
	logger: Logger = new Logger(PrismaService.name);

	constructor(private configService: ConfigService) {
		super({
			datasources: {
				db: {
					url: configService.get<string>('DATABASE_URL')
				}
			}
		});
	}

	async onModuleInit() {
		await this.$connect();
		this.logger.log('onModuleInit prisma connected');
	}

	async enableShutdownHooks(app: INestApplication) {
		this.$on('beforeExit', async () => {
			this.logger.log('enableShutdownHooks prisma closing');
			await app.close();
		});
	}
}
