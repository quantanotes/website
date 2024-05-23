<script>
    import { onMount } from 'svelte'

    let zoom = 1
    let clientWidth = 0
    let clientHeight = 0
    let width = 0
    let height = 0
    let cx = -0.743643887037158704752191506114774
    let cy = 0.131825904205311970493132056385139
    let text

    const fontSize = 16
    const maxIter = 1000
    const chars = [' ', '.', ',', '-', '+', '*', 'X', 'O', '#', '@']
    const _class = `h-64 overflow-clip text-[16px] leading-[16px] -z-10 w-full`
    const centers = [
        [
            -0.743643887037158704752191506114774,
            0.131825904205311970493132056385139,
        ],
        // [-1.7499462, 0],
        [-0.7746806106269039, -0.1374168856037867],
    ]

    $: width = Math.floor((clientWidth * 1.7) / fontSize)
    $: height = Math.floor(clientHeight / fontSize)

    function mandelbrot(x0, y0) {
        let x = 0.0
        let y = 0.0
        let i = 0
        let x2 = 0
        let y2 = 0
        const escape = 4

        while (x2 + y2 <= escape && i < maxIter) {
            const xtemp = x2 - y2 + x0
            y = 2 * x * y + y0
            x = xtemp
            x2 = x * x
            y2 = y * y
            i++
        }

        return i
    }

    function render(zoom) {
        let ascii = ''
        let first = 'N'
        let homo = true

        for (let y = 0; y < height; y++) {
            for (let x = 0; x < width; x++) {
                const real = ((x / width) * 3.5 - 2.5) / zoom + cx
                const imaginary = ((y / height) * 2 - 1) / zoom + cy
                const iter = mandelbrot(real, imaginary)
                const index = Math.floor(iter) % chars.length
                const char = chars[index]

                if (first == 'N') {
                    first = char
                } else if (first != char) {
                    homo = false
                }

                ascii += char
            }

            ascii += '\n'
        }

        if (homo) {
            reset()
        }

        return ascii
    }

    function reset() {
        const center = centers[Math.floor(Math.random() * centers.length)]
        zoom = 1
        cx = center[0]
        cy = center[1]
    }

    onMount(() => {
        reset()

        function animate() {
            zoom *= 1.02
            text = render(zoom)
        }

        const interval = setInterval(animate, 20)

        return () => clearInterval(interval)
    })
</script>

<div class={_class} bind:clientWidth bind:clientHeight>
    <pre class="text-contrast">{text}</pre>
</div>
