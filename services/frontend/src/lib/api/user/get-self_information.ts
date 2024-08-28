import { env } from "$env/dynamic/private";
import ky from "ky";
import type { SelfInformation } from "./models/SelfInformation";

export async function getSelfInformation(jwt: string): Promise<SelfInformation | null> {
    return (await ky.get(env.BACKEND_URL + 'api/v1/self', { headers: { "Authorization": jwt } }).json<SelfInformation>());
}