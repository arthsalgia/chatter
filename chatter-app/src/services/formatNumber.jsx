export default function formatNumber(num) {
    const strNum = num ? String(num) : "";
    if (strNum.startsWith("+")) { 
        const lastPart = strNum.slice(-10);
        const firstPart = strNum.slice(0, strNum.length - 10);
        return firstPart + " " + lastPart;
    }
    return strNum;
}