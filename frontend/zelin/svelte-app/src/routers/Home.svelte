<script>
    import HeadBar from "./bars/head_bar.svelte";
    import TitleBar from "./bars/title_bar.svelte";
    import AppLeftBar from "./bars/app_left_bar.svelte";
    import RightColumn from "./bars/right_column.svelte";
    import MidColumn from "./bars/mid_column.svelte";
</script>

<main>
    <HeadBar />
    <TitleBar />
    <AppLeftBar />
    <RightColumn />
    <MidColumn />
</main>

<style>
    * {
	  box-sizing: border-box;
	}
</style>