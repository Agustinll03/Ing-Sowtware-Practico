import React from "react";
import { useEffect, useState } from "react";

export default function App() {
  const [items, setItems] = useState([]);
  const [name, setName] = useState("");

  useEffect(() => {
    fetch("/api/items")
      .then(res => res.json())
      .then(setItems);
  }, []);

  const addItem = async (e) => {
    e.preventDefault();
    await fetch("/api/items", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ name }),
    });
    setName("");
    const res = await fetch("/api/items");
    setItems(await res.json());
  };

  return (
    <div>
      <h1>Items</h1>
      <form onSubmit={addItem}>
        <input value={name} onChange={(e) => setName(e.target.value)} />
        <button>Add</button>
      </form>
      <ul>
        {items.map(it => <li key={it.id}>{it.name}</li>)}
      </ul>
    </div>
  );
}