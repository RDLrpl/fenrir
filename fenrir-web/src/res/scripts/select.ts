import { defineComponent, ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';

export default defineComponent({
  setup() {
    const router = useRouter();
    const routeElement = ref<HTMLElement | null>(null);

    const toTG = () => {
        router.push('/tg');
    };

    const toSettings = () => {
        router.push('/settings');
    };

    const setRandomValue = () => {
        if (routeElement.value) {
            const randomNumber: number = Math.floor(Math.random() * 21) - 10;
            routeElement.value.style.setProperty('--rnd-i', randomNumber.toString());
        }
    };

    onMounted(() => {
        setRandomValue();
    });

    return {
        defineComponent,
        toTG,
        toSettings,
        setRandomValue
    };
  }
});