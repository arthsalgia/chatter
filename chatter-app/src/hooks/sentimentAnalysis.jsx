export default async function sentimentAnalysisApi(chat) {
  if (chat == "") {
    return
  }

  const res = await fetch(
    `http://localhost:8000/api/sentiment-analysis?chat=${encodeURIComponent(chat)}`,
    {
      method: "GET",
    }
  );

  if (!res.ok) {
    throw new Error("Failed to get sentiment analysis api");
  }

  return await res.json();
}