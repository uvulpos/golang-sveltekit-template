import { redirect } from "@sveltejs/kit";

/** @type {import('./$types').LayoutServerLoad} */
export async function load({ cookies }) {
    const jwt = cookies.get('jwt');

    console.log("FRONTEND JWT: ", jwt);


    if (jwt === undefined || jwt === "") {
        throw redirect(301, "/login");
    }
    return {
        status: 200,
    };
}