import { checkUserPermissionMiddleware } from "$lib/functions/permissions/check-user-permission-middleware";
import { error } from "@sveltejs/kit";

/** @type {import('./$types').LayoutServerLoad} */
export async function load({ cookies, fetch }) {
	try {
		const jwt = cookies.get('jwt');

		if (jwt == null || jwt == "") {
			error(401, {
				message: "Unauthorized",
			});
		}

		const value = await checkUserPermissionMiddleware(jwt, ["admin"]);
		return {
			status: 200,
			props: {
				value,
			}
		};
	} catch (x) {
		error(403, {
			message: "Forbidden",
		});
	}
}