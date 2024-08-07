
/** @type {import('./$types').LayoutServerLoad} */
export async function load({ cookies }) {
	const jwt = cookies.get('jwt');
	console.log('JWT:', jwt);

	if (jwt == null || jwt == "") {
		throw new NotPermittedError();
	}
	return {};
}

class NotPermittedError extends Error {
	constructor() {
		const message = "User not Authorized"
		super(message);
		this.name = 'NotPermittedError';
		this.message = "You do not have enough permissions";
	}
}