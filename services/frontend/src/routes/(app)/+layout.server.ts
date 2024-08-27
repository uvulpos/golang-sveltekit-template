import { getSelfInformation } from "$lib/api/user/get-self_information";
import { redirect } from "@sveltejs/kit";


/** @type {import('./$types').LayoutServerLoad} */
export async function load({ cookies, fetch }) {
    const jwt = cookies.get('jwt');

    if (jwt === undefined || jwt === "") {
        throw redirect(301, "/login");
    }

    let self = await getSelfInformation(fetch, jwt ?? "");

    if (self === undefined || self == null) {
        throw redirect(301, "/login");
    }

    self.profile_picture = "https://avatars.githubusercontent.com/u/53957681?v=4";

    return {
        props: {
            user: self,
        },
        status: 200,
    };
}