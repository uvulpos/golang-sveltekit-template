import { checkUserPermissionMiddleware, NotPermittedError } from "$lib/functions/permissions/check-user-permission-middleware";

/** @type {import('./$types').LayoutServerLoad} */
export async function load({ cookies, fetch }) {
	try {
		const jwt = cookies.get('jwt');

		if (jwt == null || jwt == "") {
			throw new NotPermittedError();
		}

		const value = await checkUserPermissionMiddleware(fetch, jwt, ["admin-super"]);
		return {
			status: 200,
			props: {
				value,
			}
		};
	} catch (error) {
		return {
			status: 401,
			error: error instanceof NotPermittedError ? 'NotPermittedError' : 'UnknownError'
		};
	}
}