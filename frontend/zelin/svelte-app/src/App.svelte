
<script>

	import Home from './routers/Home.svelte';
	import Backend from './routers/backend/Backend.svelte';
	import router from 'page';

	let page
	
	router('/', () => page = Home)
	router('/backend', () => page = Backend)
	router.start()
</script>

<main>
	<svelte:component this={page} />
</main>

<style>

</style>