export default async function chats() {

  const res = await fetch(
    `http://localhost:8000/api/get-all-chats`,
    {
      method: "GET",
    }
  );

  if (!res.ok) {
    throw new Error("Failed to get chats api");
  }

  return await res.json();
}