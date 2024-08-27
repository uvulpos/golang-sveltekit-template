import { redirect } from "@sveltejs/kit";

/** @type {import('./$types').LayoutServerLoad} */
export async function load({ cookies }) {
    // cookies.delete("jwt", { path: "/" });
    throw redirect(301, "/api/v1/oauth/logout");
}