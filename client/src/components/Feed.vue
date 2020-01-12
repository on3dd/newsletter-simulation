<template>
    <b-container class="p-2">
        <b-row class="">
            <b-col>
                <b-container>
                    <h1 style="margin: 0">Feed</h1>
                </b-container>
            </b-col>
        </b-row>
        <b-row class="my-4">
            <b-col>
                <b-container fluid>
                    <b-button
                            v-if="unseen.length != 0"
                            variant="outline-primary"
                            class="update-posts-btn"
                            size="lg"
                            @click="renderUnseen"
                    >
                        {{ unseen.length }} new {{ unseen.length == 1 ? 'post' : 'posts' }}
                    </b-button>
                </b-container>
            </b-col>
        </b-row>
        <b-row
                v-for="(item, index) in orderedPosts"
                :key="index"
                class="post"
        >
            <b-col>
                <NewsItem
                        :id="item.id"
                        :author="item.author"
                        :posted_at="timeToDate(item.posted_at)"
                        :content="item.content"
                />
                <hr>
            </b-col>
        </b-row>
    </b-container>
</template>

<script>
    import NewsItem from "./NewsItem";

    export default {
        name: "Feed",
        components: {
            NewsItem
        },
        data() {
            return {
                logs: [],
                unseen: [],
                posts: [
                    {
                        id: 1,
                        author: {
                            id: 1,
                            name: "Barack Obama",
                            image: "https://avatars0.githubusercontent.com/u/45851782?s=460&v="
                        },
                        posted_at: "2020-01-12T11:51:20.360547726+10:00",
                        content: {
                            text: "Lorem ipsum dolor sit amet, consectetur adipisicing elit. A debitis est iste nulla quos. Amet consequuntur, cupiditate eaque illo incidunt libero necessitatibus odio, omnis pariatur provident quas quod tempore vitae.",
                            image: "https://picsum.photos/1920/1080"
                        }
                    },
                    {
                        id: 2,
                        author: {
                            id: 2,
                            name: "Vova Putin",
                            image: "https://deita.ru/media/images/___________NVxAkAP.2e16d0ba.fill-950x690-c100.jpg"
                        },
                        posted_at: "2020-01-12T11:52:40.360547726+10:00",
                        content: {
                            text: "Lorem ipsum dolor sit amet, consectetur adipisicing elit. A debitis est iste nulla quos. Amet consequuntur, cupiditate eaque illo incidunt libero necessitatibus odio, omnis pariatur provident quas quod tempore vitae.",
                            image: "https://picsum.photos/1920/1000"
                        }
                    }
                ]
            }
        },
        created() {
            this.connect();
        },
        methods: {
            connect() {
                this.socket = new WebSocket("ws://localhost:8080/ws");

                this.socket.onopen = () => {
                    this.logs.push({event: 'WebSocket Connected', data: this.socket.url});

                    this.socket.onmessage = ({data}) => {
                        this.logs.push({event: 'Recieved message', data});
                        this.unseen.push(JSON.parse(data))
                    }
                }
            },
            renderUnseen() {
                this.posts = this.posts.concat(this.unseen)
                this.unseen = []
            },
            timeToDate(time) {
                const options = {
                    year: 'numeric',
                    month: 'long',
                    day: 'numeric',
                    timezone: 'UTC',
                    hour: 'numeric',
                    minute: 'numeric',
                    second: 'numeric'
                }

                return new Date(Date.parse(time)).toLocaleDateString("en-US", options)
            }
        },
        computed: {
            orderedPosts() {
                return _.orderBy(this.posts, "posted_at", "desc");
            },
        }
    }

</script>

<style scoped lang="scss">
    .update-posts-btn {
        width: 100%;
    }
    hr {
        margin-top: 2rem;
        margin-bottom: 2rem;
        border: 0;
        border-top: 1px solid rgba(0, 0, 0, 0.1);
    }
</style>