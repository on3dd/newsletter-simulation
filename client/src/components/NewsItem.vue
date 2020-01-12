<template>
    <b-container fluid class="post">
        <b-row class="header">
            <b-col>
                <div class="d-flex">
                    <div class="author-img mr-2">
                        <a :href="'users/'+id">
                            <img :src="authorImg" alt="">
                        </a>
                    </div>
                    <div class="info">
                        <div class="d-flex flex-column p-1">
                            <a :href="'users/'+id" style="line-height: 20px">
                                <span>{{ authorName }}</span>
                            </a>
                            <small :title="posted_at">{{ posted_at }}</small>
                        </div>
                    </div>
                </div>
            </b-col>
        </b-row>
        <b-row class="mt-2">
            <b-col>
                <p class="content-text">
                    {{ content.text }}
                </p>
            </b-col>
        </b-row>
        <b-row>
            <b-col>
                <b-container fluid class="content-img">
                    <img :src="contentImg" alt="">
                </b-container>
            </b-col>
        </b-row>
    </b-container>
</template>

<script>
    export default {
        name: "NewsItem",
        props: {
            id: {
                required: true
            },
            author: {
                required: true
            },
            posted_at: {
                required: true
            },
            content: {
                required: true
            }
        },
        computed: {
            authorName() {
                return this.author.name
            },
            authorImg() {
                return this.author.image
            },
            contentImg() {
                return this.content.image
            }
        }
    }
</script>

<style scoped lang="scss">
    .post {
        .author-img {
            width: 50px;
            height: 50px;
            position: relative;

            img {
                position: absolute;
                top: 50%;
                left: 50%;
                width: 100%;
                height: 100%;
                transform: translate(-50%, -50%);
                border-radius: 50%;
            }
        }

        .info {
            small {
                cursor: pointer;
                &:hover {
                    text-decoration: underline;
                }
            }
        }

        .content-img {
            overflow: hidden;
            position: relative;
            height: 400px;
            border-radius: 20px;

            img {
                position: absolute;
                top: 50%;
                left: 50%;
                transform: translate(-50%, -50%);
            }
        }
    }
</style>