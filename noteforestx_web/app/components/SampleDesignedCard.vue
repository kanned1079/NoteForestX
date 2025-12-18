<script setup lang="ts">

</script>

<template>
  <div class="nhost-card">
    <div class="card-glow"></div>
    <div class="card-grid"></div>

    <div class="card-header">
      <div class="brand-wrapper">
        <div class="logo-container">
          <svg
              class="nhost-logo"
              viewBox="0 0 32 34"
              fill="none"
              xmlns="http://www.w3.org/2000/svg"
              aria-label="Nhost Logo"
          >
            <path
                d="M29.013 7.305 17.12.46a3.489 3.489 0 0 0-3.461 0 3.458 3.458 0 0 0-1.729 2.984v.894l-.775-.447a3.489 3.489 0 0 0-3.461 0 3.459 3.459 0 0 0-1.729 2.99v.892l-.775-.446a3.488 3.488 0 0 0-3.46 0A3.459 3.459 0 0 0 0 10.313v21.443a1.624 1.624 0 0 0 1.803 1.612c.303-.033.591-.15.83-.34l5.898-4.636 9.096 5.233a1.652 1.652 0 0 0 1.627 0c.501-.29.814-.826.814-1.404V19.319a5.952 5.952 0 0 0-2.983-5.149l-2.983-1.716V3.448a1.288 1.288 0 0 1 1.291-1.286c.227 0 .449.06.645.173L27.93 9.177a3.795 3.795 0 0 1 1.898 3.276V28.52c0 .458-.248.884-.645 1.113l-3.151 1.813V15.884a5.953 5.953 0 0 0-2.983-5.148L15.73 6.524V9.02l6.237 3.59a3.793 3.793 0 0 1 1.898 3.275V32.38c0 .575.311 1.115.813 1.405a1.653 1.653 0 0 0 1.628 0l3.965-2.282A3.457 3.457 0 0 0 32 28.517V12.448a5.968 5.968 0 0 0-2.987-5.143ZM15.995 16.04a3.792 3.792 0 0 1 1.898 3.277v11.97l-7.53-4.333 2.417-1.898a3.424 3.424 0 0 0 1.318-2.707v-7.397l1.898 1.09v-.002Zm-4.066-2.34v8.645c0 .397-.18.766-.491 1.01l-9.27 7.283V10.31a1.286 1.286 0 0 1 1.29-1.286c.226 0 .448.06.644.173l1.863 1.07v15.325l2.168-1.704V6.878a1.286 1.286 0 0 1 1.29-1.285c.227 0 .45.06.645.173l1.861 1.07v4.367L9.762 9.955v2.498l2.17 1.249h-.003Z"
                fill="currentColor"
            ></path>
          </svg>
        </div>
        <span class="brand-text">NHOST</span>
      </div>

      <div class="action-buttons">
        <button class="btn-icon" aria-label="Star this repo">
          <svg
              class="icon star"
              viewBox="0 0 16 16"
              xmlns="http://www.w3.org/2000/svg"
          >
            <path
                fill-rule="evenodd"
                d="M8 .25l2.472 5.008 5.528.803-4 3.899.944 5.508L8 12.625l-4.944 2.843.944-5.508-4-3.9 5.528-.802L8 .25z"
            ></path>
          </svg>
        </button>
        <button class="btn-icon" aria-label="View on GitHub">
          <svg
              class="icon github"
              viewBox="0 0 16 16"
              version="1.1"
              aria-hidden="true"
          >
            <path
                fill-rule="evenodd"
                d="M8 0C3.58 0 0 3.58 0 8c0 3.54 2.29 6.53 5.47 7.59.4.07.55-.17.55-.38 0-.19-.01-.82-.01-1.49-2.01.37-2.53-.49-2.69-.94-.09-.23-.48-.94-.82-1.13-.28-.15-.68-.52-.01-.53.63-.01 1.08.58 1.23.82.72 1.21 1.87.87 2.33.66.07-.52.28-.87.51-1.07-1.78-.2-3.64-.89-3.64-3.95 0-.87.31-1.59.82-2.15-.08-.2-.36-1.02.08-2.12 0 0 .67-.21 2.2.82.64-.18 1.32-.27 2-.27.68 0 1.36.09 2 .27 1.53-1.04 2.2-.82 2.2-.82.44 1.1.16 1.92.08 2.12.51.56.82 1.27.82 2.15 0 3.07-1.87 3.75-3.65 3.95.29.25.54.73.54 1.48 0 1.07-.01 1.93-.01 2.2 0 .21.15.46.55.38A8.013 8.013 0 0016 8c0-4.42-3.58-8-8-8z"
            ></path>
          </svg>
        </button>
      </div>
    </div>

    <div class="card-body">
      <div class="repo-title">nhost-js<span class="blinking-cursor"></span></div>
      <p class="repo-description">
        The ultra-lightweight JavaScript SDK for interacting with the Nhost stack.
      </p>
      <div class="tag-wrapper">
        <span class="badge badge-js">JS</span>
        <span class="badge badge-ts">TS</span>
      </div>
    </div>
  </div>

</template>

<style scoped>
@keyframes blink {
  0%,
  100% {
    opacity: 1;
  }
  50% {
    opacity: 0;
  }
}

@keyframes grid-move {
  0% {
    background-position: 0 0;
  }
  100% {
    background-position: 30px 30px;
  }
}

@keyframes float {
  0%,
  100% {
    transform: translateY(0);
  }
  50% {
    transform: translateY(-4px);
  }
}

/* Component Styles */
.nhost-card {
  --bg-color: #0c111d;
  --card-border: #1f2937;
  --text-primary: #f3f4f6;
  --text-secondary: #9ca3af;
  --accent-color: #0052cc;
  --accent-glow: rgba(0, 82, 204, 0.35);
  --brand-gradient: linear-gradient(135deg, #0052cc, #2684ff);

  font-family: "Inter", system-ui, sans-serif;
  width: 100%;
  max-width: 400px;
  background-color: var(--bg-color);
  border: 1px solid var(--card-border);
  border-radius: 16px;
  padding: 24px;
  position: relative;
  overflow: hidden;
  box-shadow: 0 4px 24px -1px rgba(0, 0, 0, 0.2);
  transition:
      transform 0.3s cubic-bezier(0.34, 1.56, 0.64, 1),
      box-shadow 0.3s ease,
      border-color 0.3s ease;
  z-index: 1;
}

.nhost-card:hover {
  transform: translateY(-4px) scale(1.01);
  box-shadow:
      0 20px 40px -5px rgba(0, 0, 0, 0.4),
      0 0 0 1px rgba(38, 132, 255, 0.3);
  border-color: rgba(38, 132, 255, 0.5);
}

.card-grid {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-image: linear-gradient(
      rgba(255, 255, 255, 0.05) 1px,
      transparent 1px
  ),
  linear-gradient(90deg, rgba(255, 255, 255, 0.05) 1px, transparent 1px);
  background-size: 30px 30px;
  mask-image: linear-gradient(to bottom, black 40%, transparent 100%);
  pointer-events: none;
  z-index: -2;
  transition: opacity 0.5s ease;
  opacity: 0.5;
}

.nhost-card:hover .card-grid {
  animation: grid-move 20s linear infinite;
  opacity: 0.8;
}

.card-glow {
  position: absolute;
  top: -50px;
  right: -50px;
  width: 150px;
  height: 150px;
  background: radial-gradient(
      circle,
      var(--accent-glow) 0%,
      rgba(0, 0, 0, 0) 70%
  );
  filter: blur(40px);
  z-index: -1;
  transition: all 0.5s ease;
}

.nhost-card:hover .card-glow {
  transform: scale(1.5);
  background: radial-gradient(
      circle,
      rgba(38, 132, 255, 0.4) 0%,
      rgba(0, 0, 0, 0) 70%
  );
}

/* Header Section */
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
}

.brand-wrapper {
  display: flex;
  align-items: center;
  gap: 12px;
}

.logo-container {
  width: 40px;
  height: 40px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: rgba(255, 255, 255, 0.03);
  border-radius: 8px;
  border: 1px solid rgba(255, 255, 255, 0.1);
  transition: background 0.3s ease;
}

.nhost-logo {
  width: 24px;
  height: 24px;
  color: #fff;
  filter: drop-shadow(0 0 2px rgba(255, 255, 255, 0.3));
}

.brand-text {
  font-size: 18px;
  font-weight: 700;
  letter-spacing: 0.5px;
  color: var(--text-primary);
}

/* Actions */
.action-buttons {
  display: flex;
  gap: 8px;
}

.btn-icon {
  background: transparent;
  border: none;
  cursor: pointer;
  padding: 8px;
  border-radius: 8px;
  color: var(--text-secondary);
  transition: all 0.2s ease;
  display: flex;
  align-items: center;
  justify-content: center;
}

.btn-icon:hover {
  background: rgba(255, 255, 255, 0.1);
  color: #fff;
  transform: translateY(-2px);
}

.icon {
  width: 20px;
  height: 20px;
  fill: currentColor;
}

.icon.star {
  transition: color 0.2s;
}

.btn-icon:hover .icon.star {
  color: #fbbf24;
  filter: drop-shadow(0 0 6px rgba(251, 191, 36, 0.5));
}

/* Body Section */
.card-body {
  position: relative;
}

.repo-title {
  font-family: "JetBrains Mono", monospace;
  font-size: 24px;
  font-weight: 500;
  color: var(--text-primary);
  margin: 0 0 12px 0;
  display: flex;
  align-items: center;
}

.blinking-cursor {
  display: inline-block;
  width: 10px;
  height: 3px;
  background-color: var(--accent-color);
  margin-left: 6px;
  animation: blink 1s step-end infinite;
  box-shadow: 0 0 8px var(--accent-color);
}

.repo-description {
  font-size: 15px;
  line-height: 1.6;
  color: var(--text-secondary);
  margin: 0 0 24px 0;
}

/* Tags */
.tag-wrapper {
  display: flex;
  gap: 10px;
}

.badge {
  font-family: "JetBrains Mono", monospace;
  font-size: 12px;
  font-weight: 600;
  padding: 4px 10px;
  border-radius: 6px;
  transition: all 0.3s ease;
  cursor: default;
  position: relative;
  overflow: hidden;
}

/* JS Tag */
.badge-js {
  background-color: rgba(247, 223, 30, 0.1);
  color: #f7df1e;
  border: 1px solid rgba(247, 223, 30, 0.2);
}

.badge-js:hover {
  background-color: rgba(247, 223, 30, 0.2);
  box-shadow: 0 0 12px rgba(247, 223, 30, 0.2);
  transform: translateY(-1px);
}

/* TS Tag */
.badge-ts {
  background-color: rgba(49, 120, 198, 0.1);
  color: #3178c6;
  border: 1px solid rgba(49, 120, 198, 0.2);
}

.badge-ts:hover {
  background-color: rgba(49, 120, 198, 0.2);
  box-shadow: 0 0 12px rgba(49, 120, 198, 0.3);
  transform: translateY(-1px);
}

</style>