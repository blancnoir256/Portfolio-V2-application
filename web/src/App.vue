<script setup lang="ts"></script>

<template>
  <main class="page">
    <section class="hero">
      <p class="badge">Portfolio v1</p>
      <h1 class="title">BN256 ポートフォリオサイト</h1>
      <p class="lead">
        開発再開しました。<br />
        このバージョンからはCI/CDを導入し、argoCDでのデプロイを行っています。
      </p>

      <div class="card">
        <p class="card-text">ソースコードは以下のgithubリポジトリから確認できます。</p>
        <a
          class="cta"
          href="https://github.com/blancnoir256/Portfolio-V2-application"
          target="_blank"
          rel="noopener noreferrer"
        >
          githubリポジトリを見る →
        </a>
      </div>
    </section>
  </main>
</template>

<style scoped>
:global(body) {
  margin: 0;
  font-family: "Inter", system-ui, -apple-system, "Segoe UI", sans-serif;
  background: radial-gradient(circle at top, #1e1b4b 0%, #0f172a 60%);
  color: #e2e8f0;
}

.page {
  min-height: 100vh;
  display: grid;
  place-items: center;
  padding: 3rem 1.5rem;
}

.hero {
  max-width: 720px;
  text-align: left;
}

.badge {
  display: inline-flex;
  align-items: center;
  gap: 0.5rem;
  background: rgba(99, 102, 241, 0.15);
  color: #a5b4fc;
  border: 1px solid rgba(99, 102, 241, 0.4);
  padding: 0.35rem 0.75rem;
  border-radius: 999px;
  font-size: 0.85rem;
  letter-spacing: 0.06em;
  text-transform: uppercase;
}

.title {
  font-size: clamp(2.2rem, 4vw, 3.2rem);
  margin: 1rem 0 0.75rem;
  line-height: 1.1;
}

.lead {
  font-size: 1.05rem;
  color: #cbd5f5;
  margin-bottom: 1.5rem;
}

.card {
  background: rgba(15, 23, 42, 0.6);
  border: 1px solid rgba(148, 163, 184, 0.2);
  box-shadow: 0 20px 50px rgba(0, 0, 0, 0.35);
  border-radius: 16px;
  padding: 1.25rem 1.5rem;
  backdrop-filter: blur(10px);
}

.card-text {
  margin: 0 0 0.75rem;
  color: #e2e8f0;
}

.cta {
  display: inline-flex;
  align-items: center;
  gap: 0.4rem;
  padding: 0.65rem 1rem;
  border-radius: 10px;
  background: linear-gradient(135deg, #6366f1, #3b82f6);
  color: white;
  text-decoration: none;
  font-weight: 600;
  transition: transform 0.15s ease, box-shadow 0.15s ease;
}

.cta:hover {
  transform: translateY(-2px);
  box-shadow: 0 10px 25px rgba(59, 130, 246, 0.35);
}
</style>
