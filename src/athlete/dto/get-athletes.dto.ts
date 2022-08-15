import { Expose, Type } from 'class-transformer';
import { IsEnum, IsInt, IsNotEmpty, IsString, Max, Min } from 'class-validator';
import Default from 'src/common/decorators/validator.decorator';

export class GetAthletesDTO {
	@Expose()
	@Default(50)
	@Type(() => Number)
	@IsInt()
	@Min(1)
	@Max(50)
	readonly limit: number;

	@Expose()
	@Default(1)
	@Type(() => Number)
	@IsInt()
	@Min(1)
	readonly page: number;

	@Expose()
	@Default('id')
	@IsString()
	@IsNotEmpty()
	readonly order_by: string;

	@Expose()
	@Default('asc')
	@IsEnum(['asc', 'desc'])
	readonly order: 'asc' | 'desc';
}
