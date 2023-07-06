import { userProfileStore, type userProfileType } from "./userProfile";

export function setUserProfile(user: userProfileType | undefined) {
    userProfileStore.set(user)
}