import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import { ConfigService } from '@nestjs/config';
import helmet from 'helmet';
import { ValidationPipe } from '@nestjs/common';
import { Logger } from 'nestjs-pino';
import { PrismaService } from './prisma/prisma.service';

async function bootstrap() {
	const app = await NestFactory.create(AppModule);
	const config = app.get<ConfigService>(ConfigService);

	app.setGlobalPrefix('v1');
	app.useLogger(app.get<Logger>(Logger));
	app.useGlobalPipes(new ValidationPipe({ transform: true }));
	app.enableCors();
	app.use(helmet());

	const prismaService: PrismaService = app.get(PrismaService);
	prismaService.enableShutdownHooks(app);

	await app.listen(config.get<number>('PORT') as number);
}
bootstrap();
