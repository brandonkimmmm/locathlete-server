export const getPaginationQuery = (limit: number, page: number) => {
	return {
		take: limit,
		skip: limit * (page - 1)
	};
};

export const getOrderingQuery = (orderBy: string, order: 'asc' | 'desc') => {
	return {
		orderBy: [
			{
				[orderBy]: order
			}
		]
	};
};
