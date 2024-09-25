import { Title } from "@solidjs/meta";
import { render } from "solid-js/web";
import { createSignal } from "solid-js";
import { marked } from "marked";

const t = () => {
  const el = <div></div>
  el.innerHTML = marked.parse('# Marked in the browser\n\nRendered by **marked**.');
  return el;
}

export default function Home() {
  const [article, setArticle] = createSignal(t)

  return (
    <main>
      <Title>Couryrr - Blog</Title>
      <h1 class="text-3xl font-bold underline">Blog</h1>
      {article()}
    </main>
  );
}
