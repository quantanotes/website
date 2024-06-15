<script>
    import { inview } from 'svelte-inview'
    import SectionDivider from '$/components/atoms/section-divider.svelte'
    import Thread from '../thread.svelte'

    const usecases = [
        {
            name: 'life coach',
            thread: [
                {
                    role: 'user',
                    content: "I feel stuck in my career and don't know what to do next.",
                },
                {
                    role: 'assistant',
                    content:
                        'I understand how challenging that can be. Let’s explore your skills and passions together to find the right path for you. Would you like to start with a strengths assessment or discuss your interests and values?',
                },
            ],
        },
        {
            name: 'software engineer',
            thread: [
                {
                    role: 'user',
                    content:
                        'I need help debugging this code. It keeps throwing an undefined error.',
                },
                {
                    role: 'assistant',
                    content:
                        'I can help with that! Please share your code snippet and I’ll analyze it for potential issues. Additionally, I can provide documentation or examples that might help clarify the problem.',
                },
            ],
        },
        {
            name: 'personal trainer',
            thread: [
                {
                    role: 'user',
                    content: 'I want to start a new workout routine but don’t know where to begin.',
                },
                {
                    role: 'assistant',
                    content:
                        'Starting a new workout routine is a great step! I can create a personalized plan based on your fitness level and goals. Would you like to focus on strength, cardio, flexibility, or a combination of these?',
                },
            ],
        },
        {
            name: 'medical advisor',
            thread: [
                {
                    role: 'user',
                    content: 'I’ve been feeling fatigued lately. What could be the cause?',
                },
                {
                    role: 'assistant',
                    content:
                        'Fatigue can have many causes, from lifestyle factors to medical conditions. I can help you narrow it down by asking some questions about your sleep, diet, stress levels, and activity. Would you like to proceed?',
                },
            ],
        },
        {
            name: 'sales automation',
            thread: [
                {
                    role: 'user',
                    content: 'I need to streamline my sales process. Any suggestions?',
                },
                {
                    role: 'assistant',
                    content:
                        'Absolutely! I can set up automated email campaigns, manage your CRM, and analyze your sales data to identify areas for improvement. Would you like to start with automating your outreach or analyzing your current sales pipeline?',
                },
            ],
        },
        {
            name: 'public relations specialist',
            thread: [
                { role: 'user', content: 'How can I improve my company’s public image?' },
                {
                    role: 'assistant',
                    content:
                        'Improving your public image is a multi-faceted task. I can help you with crafting press releases, engaging with your audience on social media, and monitoring public perception. Where would you like to focus first?',
                },
            ],
        },
        {
            name: 'researcher',
            thread: [
                {
                    role: 'user',
                    content: 'Can you help me find the latest studies on climate change?',
                },
                {
                    role: 'assistant',
                    content:
                        'Sure, I can find and summarize the latest research on climate change for you. Do you need information on a specific aspect, such as its effects, mitigation strategies, or policy developments?',
                },
            ],
        },
        {
            name: 'virtual tutor',
            thread: [
                {
                    role: 'user',
                    content: 'I’m struggling with calculus. Can you help me understand integrals?',
                },
                {
                    role: 'assistant',
                    content:
                        'Of course! I can explain the concept of integrals, provide examples, and guide you through practice problems. Would you like to start with the basics or jump into a specific problem you’re working on?',
                },
            ],
        },
        {
            name: 'market analyst',
            thread: [
                { role: 'user', content: 'I need a market analysis for the tech industry.' },
                {
                    role: 'assistant',
                    content:
                        'I can provide a comprehensive market analysis including trends, competitor analysis, and consumer insights in the tech industry. Is there a particular segment or geographic area you want to focus on?',
                },
            ],
        },
        {
            name: 'customer service representative',
            thread: [
                { role: 'user', content: 'I need help with my order. It hasn’t arrived yet.' },
                {
                    role: 'assistant',
                    content:
                        'I’m sorry to hear that your order hasn’t arrived. Let me check the status for you. Can you please provide your order number?',
                },
            ],
        },
        {
            name: 'event planner',
            thread: [
                {
                    role: 'user',
                    content: 'I’m planning a corporate event and need help with the details.',
                },
                {
                    role: 'assistant',
                    content:
                        'I’d love to help you plan your corporate event! I can assist with venue selection, catering, entertainment, and more. What’s the theme or objective of your event?',
                },
            ],
        },
        {
            name: 'roleplaying',
            thread: [
                { role: 'user', content: 'I want to create a fantasy adventure storyline.' },
                {
                    role: 'assistant',
                    content:
                        'That sounds like fun! Let’s start by building your world. Do you have a specific setting in mind, or should we brainstorm some ideas together?',
                },
            ],
        },
        {
            name: 'legal advisor',
            thread: [
                {
                    role: 'user',
                    content: 'I need advice on creating a contract for freelance work.',
                },
                {
                    role: 'assistant',
                    content:
                        'I can help you draft a comprehensive contract for your freelance work. Let’s go over the essential elements such as scope of work, payment terms, and deadlines. What specific details do you need included?',
                },
            ],
        },
        {
            name: 'personal assistant',
            thread: [
                { role: 'user', content: 'Can you manage my calendar for the week?' },
                {
                    role: 'assistant',
                    content:
                        'Absolutely, I can organize your calendar, schedule appointments, and set reminders for you. Please share your current schedule and any specific tasks or events you need to prioritize.',
                },
            ],
        },
        {
            name: 'internet search engine',
            thread: [
                {
                    role: 'user',
                    content:
                        'Can you find me the best-reviewed Italian restaurants in New York City?',
                },
                {
                    role: 'assistant',
                    content:
                        'Sure! I’ll compile a list of the best-reviewed Italian restaurants in New York City, including their ratings, reviews, and addresses. Please hold on for a moment.',
                },
            ],
        },
        {
            name: 'web scraping',
            thread: [
                { role: 'user', content: 'I need data from this website for my project.' },
                {
                    role: 'assistant',
                    content:
                        'I can scrape the data you need from the website. Please provide the URL and specify the type of data you’re looking for. I’ll handle the rest!',
                },
            ],
        },
        {
            name: 'cold emailer',
            thread: [
                {
                    role: 'user',
                    content: 'I need help drafting cold emails for potential clients.',
                },
                {
                    role: 'assistant',
                    content:
                        'I can craft engaging cold emails tailored to your target audience. Could you provide some details about your services and the type of clients you’re reaching out to?',
                },
            ],
        },
        {
            name: 'content planning',
            thread: [
                { role: 'user', content: 'I need to plan content for my blog for the next month.' },
                {
                    role: 'assistant',
                    content:
                        'Great! I can help you develop a content calendar, brainstorm topics, and even draft outlines. What’s the main theme or niche of your blog?',
                },
            ],
        },
        {
            name: 'accountability partner',
            thread: [
                { role: 'user', content: 'I want to stay on track with my fitness goals.' },
                {
                    role: 'assistant',
                    content:
                        'I can be your accountability partner and help you stay motivated. Let’s set some specific, achievable goals and I’ll check in with you regularly. How does that sound?',
                },
            ],
        },
        {
            name: 'responsive journal',
            thread: [
                {
                    role: 'user',
                    content: 'I need to reflect on my day. Can you help me with that?',
                },
                {
                    role: 'assistant',
                    content:
                        'Of course. Let’s start by focusing on the positive aspects of your day. What were three things you’re grateful for today?',
                },
            ],
        },
        {
            name: 'meditation guru',
            thread: [
                { role: 'user', content: 'I’m feeling stressed and need to relax.' },
                {
                    role: 'assistant',
                    content:
                        'I can guide you through a meditation session to help you relax. Would you prefer a breathing exercise, a guided visualization, or a mindfulness practice?',
                },
            ],
        },
        {
            name: 'story telling',
            thread: [
                { role: 'user', content: 'Tell me a bedtime story.' },
                {
                    role: 'assistant',
                    content:
                        'Sure, I’d love to. Once upon a time in a magical forest, there was a little fox named Fira who had a special talent for finding hidden treasures...',
                },
            ],
        },
        {
            name: 'copywriting',
            thread: [
                { role: 'user', content: 'I need compelling copy for my product launch.' },
                {
                    role: 'assistant',
                    content:
                        'I can craft persuasive and engaging copy for your product launch. Please share some details about your product, target audience, and key messages you want to include',
                },
            ],
        },
    ]

    let currentUsecase = $state(Math.floor(Math.random() * usecases.length))
    let thread = $state([])
    let inView = $state(false)
    let paused = $state(false)
    let pauseTimer = $state(0)

    function animateThread() {
        if (!inView) return

        if (paused) {
            pauseTimer += 20
            if (pauseTimer > 1500) {
                pauseTimer = 0
                paused = false
                thread = []
                currentUsecase = (currentUsecase + 1) % usecases.length
            }
            return
        }

        const currentMessageIndex = thread.length - 1
        const referenceThread = usecases[currentUsecase].thread
        const referenceMessage = referenceThread[currentMessageIndex]
        const currentMessage = thread[thread.length - 1]
        const nextMessage = referenceThread[currentMessageIndex + 1]

        if (!currentMessage || currentMessage.content.length === referenceMessage.content.length) {
            if (nextMessage) {
                thread = [...thread, { role: nextMessage.role, content: '' }]
            } else {
                paused = true
            }
        } else {
            const content = referenceMessage.content.slice(0, currentMessage.content.length + 3)
            thread[thread.length - 1] = { ...currentMessage, content }
        }
    }

    setInterval(animateThread, 20)
</script>

<div class="p-4">
    <h1 class="flex font-bold justify-between mx-auto text-4xl text-center w-fit">Use cases</h1>

    <br />
    <br />
    <br />
    <br />

    <div class="flex flex-col flex-wrap h-96 max-w-4xl mx-auto">
        {#each usecases as usecase, i}
            <!-- svelte-ignore a11y_no_static_element_interactions -->
            <div
                class={`duration-500 transition-all w-fit ${i == currentUsecase && 'font-bold text-2xl'}`}
                onmouseenter={() => {
                    if (i !== currentUsecase) {
                        currentUsecase = i
                        thread = []
                        paused = false
                        pauseTimer = 0
                    }
                }}
            >
                {usecase.name}
            </div>
        {/each}
    </div>

    <br />
    <br />
    <br />
    <br />

    <div
        class="h-96 max-w-4xl mx-auto"
        use:inview
        oninview_change={(event) => {
            inView = event.detail.inView
        }}
    >
        <Thread bind:thread disableTransitions={true} />
    </div>
</div>

<SectionDivider />
