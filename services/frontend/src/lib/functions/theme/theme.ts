import { ThemeEnum } from "$lib/stores";

export function changePageTheme(bodyElement: HTMLElement | undefined, theme: ThemeEnum) {
    const bodyClassList = bodyElement?.classList;
    const dakrmodeClassName = "darkmode";

    if (bodyClassList === null || bodyClassList === undefined) {
        return;
    }

    if (theme === ThemeEnum.Lightmode) {
        if (bodyClassList.contains(dakrmodeClassName)) {
            bodyClassList.remove(dakrmodeClassName);
            return;
        }
    }
    else {
        bodyClassList.add(dakrmodeClassName);
    }
}