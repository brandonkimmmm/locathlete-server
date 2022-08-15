import { Module } from '@nestjs/common';
import { ConfigModule } from '@nestjs/config';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { PrismaModule } from './prisma/prisma.module';
import configuration from './config';
import { nanoid } from 'nanoid';
import { LoggerModule } from 'nestjs-pino';

@Module({
	imports: [
		ConfigModule.forRoot({
			load: [configuration],
			isGlobal: true
		}),
		LoggerModule.forRoot({
			pinoHttp:
				process.env.NODE_ENV !== 'production'
					? {
							level: 'trace',
							genReqId: () => nanoid(),
							transport: {
								target: 'pino-pretty',
								options: {
									colorize: true,
									levelFirst: true,
									translateTime: true
								}
							},
							customLogLevel: (res, err) => {
								if (res) {
									if (
										res.statusCode >= 400 &&
										res.statusCode < 500
									) {
										return 'warn';
									}
									if (res.statusCode >= 500 || err) {
										return 'error';
									}
								}
								return 'info';
							},
							serializers: {
								req(req) {
									return {
										id: req.id,
										method: req.method,
										url: req.url,
										host: req.headers.host,
										origin: req.headers.origin,
										remoteAddress: req.remoteAddress
									};
								},
								res(res) {
									return {
										statusCode: res.statusCode
									};
								},
								err(err) {
									return {
										type: err.type,
										message: err.message,
										stack: err.stack,
										code: err.code
									};
								}
							},
							quietReqLogger: true
					  }
					: {}
		}),
		PrismaModule
	],
	controllers: [AppController],
	providers: [AppService]
})
export class AppModule {}
