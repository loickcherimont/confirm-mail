<!-- 
    MAIN TITLE
    ModelReadme : Replace with project name
    Language : Replace with main language for the project

    Ex: Todo_ReactJS 
-->
# ConfirmationMail_Go

<!-- ![MainImageForProject](https://placehold.co/500x300) -->

<!-- 
    INTRODUCTION => ABOUT
    Short summary explaining the reasons of the project and tell about worked concepts

    Ex: This is a Front-End project made with HTML, CSS for design and JavaScript to work on client-side validation (Constraint API)
-->

## :information_source: About
This is a back-end project made with Go(lang) on mail confirmation

<!-- 
    TOOLS
    Short list of used tools with their versions

    Ex: 
    - Go 1.18
    - MySQL 8.0.29
    - Bootstrap 5.2.0-beta1
-->
## :wrench: Tools
- Go 1.20
- Tailwindcss 3.3.6


## :gear: Setup 

### Prerequisites

Create a `.env` file at the root of the project.  
Add your data in :  

```bash
# Replace the values by your own
# Don't forget to remove them before committing
SENDER_ADDRESS=your_email
SENDER_PASSWORD=your_password
```

### Guidelines

```bash
# 1. Clone the project into your Workspace
git clone https://github.com/loickcherimont/confirm-mail

# 2. Run the server
go run ./cmd/

# 3. The project is ready!
```

## :rocket: Features
- Send a styled HTML mail to user *(optional) => Hermes if not solution!*
- Store receiver address and password in database *(later with DB)*
- Write the README
- *Responsive design*
- To Fix : Display styled templates with Tailwindcss

<!-- 
    LICENSE
    Write Developer name with used license

    Ex: Made by Loïck Chérimont
        Under MIT License 
 -->

## :key: License
Made by Loïck Chérimont  
Under MIT License

