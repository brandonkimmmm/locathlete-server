export default () => ({
	PORT: parseInt(process.env.PORT, 10) || 3000,
	DATABASE_URL: process.env.DATABASE_URL
});
