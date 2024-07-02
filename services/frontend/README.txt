https://habr.com/ru/companies/piter/articles/744824/

https://feature-sliced.design/docs/get-started/overview

(
ui
model
api
)

1. App

    Layout.tsx
        header
        <outlet>
        footer

    css

    Router.tsx

2. Pages

    HomePage
        HomePage.tsx
        ui
            HeroBlock
                Heroblock.tsx
                style.scss
            FAQ
                FAQ.tsx
                style.scss

    About

    Profile
        Profile.tsx
        ui
        Settings

3. Widgets
    Header
    Footer
    Form
    Spinner
    FileTable
        FileTable.tsx
        ui
            File
                File.tsx
    AboutProduct

4. Features
    sortUsers
        index.ts
        model
            sortUsers.ts
            deleteUser.ts
    auth

5. Entities
    users
        index.ts
        users.ts
        api
            userApi.ts
            get
            delete
            changePassword
    product

6. Shared
    ui
        MoreButton
        IconDelete

        PageTitle
            <h1></h1>
            <p className = "description"></p>
        
        GreySubtitle
            <h2></h2>