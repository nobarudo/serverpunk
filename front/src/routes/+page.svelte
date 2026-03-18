<script>
    import { onMount } from "svelte";

    // Goから受け取るデータを格納する変数（初期値は平常心）
    let heat = $state(10);
    let leftEye = $state("・");
    let rightEye = $state("・");
    let mouth = $state("_");
    let noiseOpacity = $state(0);

    // 熱量(heat)に応じたグラデーション計算
    let faceColor = $derived.by(() => {
        if (heat <= 10) return "hsl(200, 100%, 40%)";
        const activeHeat = heat - 10;
        const h = 200 - activeHeat * (200 / 90);
        const l = 40 - activeHeat * (20 / 90);
        return `hsl(${h}, 100%, ${l}%)`;
    });

    // コンポーネントが画面に表示された時に実行（APIポーリング）
    onMount(() => {
        const fetchStatus = async () => {
            try {
                const res = await fetch("/serverpunk/api/status");
                if (res.ok) {
                    const data = await res.json();
                    heat = data.heat;
                    leftEye = data.left_eye;
                    rightEye = data.right_eye;
                    mouth = data.mouth;
                    noiseOpacity = data.noise / 100;
                }
            } catch (e) {
                console.error("サーバーちゃん通信エラー:", e);
            }
        };

        fetchStatus();
        const interval = setInterval(fetchStatus, 2000);

        return () => clearInterval(interval);
    });
</script>

<main class="server-monitor" style:background-color={faceColor}>
    <div class="face-container">
        <div class="face" class:stress={heat > 70}>
            <span class="eye">{leftEye}</span>
            <span class="mouth">{mouth}</span>
            <span class="eye">{rightEye}</span>
        </div>
    </div>
    <div class="noise" style:opacity={noiseOpacity}></div>
    <div class="scanlines"></div>
    <div class="vignette"></div>
</main>

<style>
    :global(body) {
        margin: 0;
        padding: 0;
        overflow: hidden;
        background-color: #000;
    }
    .server-monitor {
        width: 100vw;
        height: 100vh;
        display: flex;
        justify-content: center;
        align-items: center;
        position: relative;
        transition: background-color 0.8s ease;
    }
    .face {
        font-family: "Courier New", Courier, monospace;
        font-size: 12rem;
        font-weight: 900;
        color: #ffcc00;
        display: flex;
        align-items: center;
        gap: 2rem;
        z-index: 10;
        text-shadow:
            0 0 20px #ffcc00,
            0 0 40px #e67e22;
        user-select: none;
        transition: transform 0.2s;
    }
    .stress {
        animation: shake 0.1s infinite;
    }
    @keyframes shake {
        0% {
            transform: translate(2px, 2px);
        }
        50% {
            transform: translate(-2px, -2px);
        }
        100% {
            transform: translate(2px, -2px);
        }
    }
    .eye {
        font-size: 15rem;
        display: inline-block;
        animation: blink 4s infinite;
    }
    .mouth {
        font-size: 10rem;
    }
    @keyframes blink {
        0%,
        90%,
        100% {
            transform: scaleY(1);
        }
        95% {
            transform: scaleY(0.05);
        }
    }
    .noise {
        position: absolute;
        inset: 0;
        z-index: 15;
        pointer-events: none;
        background-image: url("https://upload.wikimedia.org/wikipedia/commons/b/b1/Visual_noise_animated.gif");
        mix-blend-mode: overlay;
        transition: opacity 0.5s ease;
    }
    .scanlines {
        position: absolute;
        inset: 0;
        background: linear-gradient(
            rgba(18, 16, 16, 0) 50%,
            rgba(0, 0, 0, 0.1) 50%
        );
        background-size: 100% 4px;
        z-index: 20;
        pointer-events: none;
    }
    .vignette {
        position: absolute;
        inset: 0;
        background: radial-gradient(
            circle,
            transparent 50%,
            rgba(0, 0, 0, 0.4) 100%
        );
        z-index: 21;
        pointer-events: none;
    }

    .face-container {
        display: flex;
        flex-direction: column;
        gap: 16px;
    }

    @media (min-width: 761px) {
        .face-container {
            flex-direction: row;
        }
    }
    @media (max-width: 760px) {
        .face-container {
            flex-direction: row;
        }
        .eye {
            font-size: 12rem;
            display: inline-block;
            animation: blink 4s infinite;
        }
        .mouth {
            font-size: 6rem;
        }
        .face {
            font-family: "Courier New", Courier, monospace;
            font-size: 12rem;
            font-weight: 900;
            color: #ffcc00;
            display: flex;
            align-items: center;
            gap: 1rem;
            z-index: 10;
            text-shadow:
                0 0 20px #ffcc00,
                0 0 40px #e67e22;
            user-select: none;
            transition: transform 0.2s;
        }
    }
</style>
