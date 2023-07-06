import type { userProfileType } from "$lib/stores/userProfile"
import type { AxiosResponse } from "axios";
import axios from "axios";

const apiRoute = "/api/v1/profile/getProfile"

export async function getUserProfile(): Promise<userProfileType> {
    try {
        const response: AxiosResponse<userProfileType> = await axios.get<userProfileType>(apiRoute, {
            withCredentials: true,
            headers: {
                'Access-Control-Allow-Origin': 'same-origin'
            }
        });
        const userData: userProfileType = response.data;
        return userData;
    } catch (error) {
        throw new Error('Error - could not fetch user');
    }
}
