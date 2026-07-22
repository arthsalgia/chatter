export default function formatNumber(num) {
    const lastPart = num.slice(-10, -1);
    const firstPart = num.slice(0, num.length -10)
    return firstPart + " " + lastPart
}