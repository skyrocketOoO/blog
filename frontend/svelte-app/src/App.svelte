
<script>

	import Home from './routers/Home.svelte';
	import Backend from './routers/backend/Backend.svelte';
	import router from 'page';
	import Datatable from './routers/datatable/datatable.svelte';

	let page
	
	router('/', () => page = Home)
	router('/backend', () => page = Backend)
	router('/datatable', () => page = Datatable)
	router.start()
</script>

<main>
	<svelte:component this={page} />
</main>

<style>

</style>