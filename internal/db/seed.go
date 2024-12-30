package db

import (
	"context"
	"fmt"
	"github.com/mujeebcodes/go-social/internal/store"
	"log"
	"math/rand"
)

var usernames = []string{
	"Nathan",
	"Laura",
	"Oscar",
	"Hazel",
	"Grace",
	"Adam",
	"Rachel",
	"Ethan",
	"Quincy",
	"Tina",
	"Paul",
	"Steve",
	"Bella",
	"Alice",
	"Kevin",
	"Zara",
	"Nina",
	"David",
	"Bob",
	"Hannah",
	"Fiona",
	"Uma",
	"Tracy",
	"Charlie",
	"Zane",
	"Penny",
	"Xander",
	"Quinn",
	"Wendy",
	"George",
	"Yara",
	"Chris",
	"Jack",
	"Ivy",
	"Diana",
	"Sam",
	"Eve",
	"Kathy",
	"Mark",
	"Mia",
	"Ruth",
	"Will",
	"Ian",
	"Victor",
	"Ursula",
	"Vince",
	"Liam",
	"Olivia",
	"Julia",
	"Frank",
}

var titles = []string{
	"Exploring the Go Programming Language",
	"Top 10 Tips for Writing Clean Code",
	"Understanding Goroutines and Concurrency",
	"The Future of Web Development",
	"Introduction to RESTful APIs",
	"Building a Microservices Architecture",
	"Mastering Data Structures in Go",
	"The Rise of AI in Everyday Life",
	"Functional Programming Explained",
	"How to Optimize Your Code for Performance",
	"Getting Started with Kubernetes",
	"An Overview of Cloud Computing",
	"The Importance of Cybersecurity",
	"Scaling Applications with Docker",
	"Deep Dive into Algorithms",
	"Best Practices for Testing Go Code",
	"Design Patterns for Scalable Systems",
	"Exploring the World of DevOps",
	"An Introduction to Blockchain Technology",
	"The Basics of Database Design",
}

var contents = []string{
	"Go is a statically typed programming language developed by Google.",
	"Clean code is essential for maintainability and collaboration.",
	"Goroutines enable concurrency in Go, making it highly efficient.",
	"The web development landscape is constantly evolving.",
	"RESTful APIs are the backbone of modern web services.",
	"Microservices offer flexibility and scalability for large systems.",
	"Understanding data structures is critical for efficient programming.",
	"AI is transforming industries like healthcare, finance, and more.",
	"Functional programming emphasizes immutability and pure functions.",
	"Code optimization can significantly improve application performance.",
	"Kubernetes simplifies container orchestration for deployment.",
	"Cloud computing provides scalable and flexible infrastructure.",
	"Cybersecurity is crucial for protecting sensitive data and systems.",
	"Docker containers help standardize application deployment.",
	"Algorithms are the foundation of problem-solving in programming.",
	"Testing ensures code reliability and reduces bugs.",
	"Design patterns are reusable solutions for common problems.",
	"DevOps bridges the gap between development and operations teams.",
	"Blockchain ensures transparency and security in transactions.",
	"Database design is a critical skill for any backend developer.",
}

var tags = []string{
	"tech",
	"golang",
	"webdev",
	"api",
	"cloud",
	"docker",
	"kubernetes",
	"devops",
	"microservices",
	"programming",
	"security",
	"blockchain",
	"ai",
	"machinelearning",
	"algorithms",
	"data",
	"testing",
	"performance",
	"designpatterns",
	"concurrency",
}

var comments = []string{
	"Great article! I learned so much from this.",
	"I disagree with your point about microservices.",
	"Can you explain more about goroutines? I'm a bit confused.",
	"Thanks for sharing! This was super helpful.",
	"This is exactly what I needed for my project. Thanks!",
	"I think you missed a key point about concurrency.",
	"Do you have any resources to dive deeper into Kubernetes?",
	"Wow, I had no idea about this topic. Great job!",
	"This article could use more examples, but it's still good.",
	"I implemented this and it works perfectly. Thank you!",
}

func Seed(store store.Storage) {
	ctx := context.Background()

	users := generateUsers(100)
	for _, user := range users {
		if err := store.Users.Create(ctx, user); err != nil {
			log.Printf("Error creating user %v: %v", user, err)
			continue
		}
	}

	posts := generatePosts(200, users)
	for _, post := range posts {
		if err := store.Posts.Create(ctx, post); err != nil {
			log.Printf("Error creating post %v: %v", post, err)
			continue
		}
	}

	comments := generateComments(500, users, posts)
	for _, comment := range comments {
		if err := store.Comments.Create(ctx, comment); err != nil {
			log.Printf("Error creating comment %v: %v", comment, err)
			continue
		}
	}

	log.Println("Seeding complete")
}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)
	usedNames := make(map[string]bool)

	for i := 0; i < num; i++ {
		var username string
		for {
			username = usernames[rand.Intn(len(usernames))] + fmt.Sprintf("%d", i)
			if !usedNames[username] {
				usedNames[username] = true
				break
			}
		}
		users[i] = &store.User{
			Username: username,
			Email:    username + "@example.com",
			Password: "123123",
		}
	}
	return users
}

func generatePosts(num int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, num)
	for i := 0; i < num; i++ {
		user := users[rand.Intn(len(users))]

		posts[i] = &store.Post{
			UserID:  user.ID,
			Title:   titles[rand.Intn(len(titles))],
			Content: contents[rand.Intn(len(contents))],
			Tags: []string{
				tags[rand.Intn(len(tags))],
				tags[rand.Intn(len(tags))],
			},
		}
	}

	return posts
}

func generateComments(num int, users []*store.User, posts []*store.Post) []*store.Comment {
	cms := make([]*store.Comment, num)
	for i := 0; i < num; i++ {
		cms[i] = &store.Comment{
			PostId:  posts[rand.Intn(len(posts))].ID,
			UserId:  users[rand.Intn(len(users))].ID,
			Content: comments[rand.Intn(len(comments))],
		}
	}

	return cms
}
