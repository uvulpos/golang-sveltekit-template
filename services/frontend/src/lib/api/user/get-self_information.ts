import ky from "ky";
import type { SelfInformation } from "./models/SelfInformation";

export async function getSelfInformation(): Promise<SelfInformation | null> {
    return (await ky.get('/api/v1/self', { credentials: 'include' }).json<SelfInformation>());
}