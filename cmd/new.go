/*
Copyright © 2024 GO NEST FRAMEWORK <gonestframework@gmail.com>
*/

package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new [app name]",
	Short: "Creates a new Go Nest application",
	Long:  "This command generates a new Go Nest application with a basic structure (main.go, go.mod, and modular architecture).",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		appName := args[0]

		fmt.Print("Enter the module name (e.g., github.com/username/repo): ")
		var moduleName string
		fmt.Scanln(&moduleName)

		createApp(appName, moduleName)
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}

// createApp scaffolds the new application structure
func createApp(appName string, moduleName string) {
	// Create the directory for the new app
	if err := os.Mkdir(appName, os.ModePerm); err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	// Navigate to the new directory
	if err := os.Chdir(appName); err != nil {
		fmt.Println("Error navigating to app directory:", err)
		return
	}

	// Initialize Go module
	cmd := exec.Command("go", "mod", "init", moduleName)
	if output, err := cmd.CombinedOutput(); err != nil {
		fmt.Println("Error initializing Go module:", err)
		fmt.Println(string(output))
		return
	}

	mainGoContentReplaced := strings.Replace(mainGoContent, "yourmodule", moduleName, -1)
	appModuleContentReplaced := strings.Replace(appModuleContent, "yourmodule", moduleName, -1)
	usersControllerContentReplaced := strings.Replace(usersControllerContent, "yourmodule", moduleName, -1)
	usersServiceContentReplaced := strings.Replace(usersServiceContent, "yourmodule", moduleName, -1)
	usersModuleContentReplaced := strings.Replace(usersModuleContent, "yourmodule", moduleName, -1)

	// Create the basic structure of the app
	createFile("main.go", mainGoContentReplaced)
	// createFile("common", "")
	createFile("domain/app_module.go", appModuleContentReplaced)
	createFile("domain/user/users_controller.go", usersControllerContentReplaced)
	createFile("domain/user/users_service.go", usersServiceContentReplaced)
	createFile("domain/user/users_module.go", usersModuleContentReplaced)
	// createFile("service", "")

	fmt.Printf("New Go Nest app '%s' created successfully!\n", appName)
}

// createFile creates a file with the given content
func createFile(path, content string) {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		fmt.Println("Error creating directories:", err)
		return
	}

	file, err := os.Create(path)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	if content != "" {
		_, err = file.WriteString(content)
		if err != nil {
			fmt.Println("Error writing to file:", err)
		}
	}
}

var mainGoContent = `package main

import (
	"context"
	"log"
	app "yourmodule/domain"

	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		app.Module,
		fx.Invoke(startServer),
	)
	app.Run()
}

func startServer(appModule *app.AppModule, lc fx.Lifecycle) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Printf("server is running on port :8080")
			go func() {
				if err := appModule.FiberApp.Listen(":8080"); err != nil {
					log.Fatalf("Failed to start server: %v", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			if err := appModule.FiberApp.Shutdown(); err != nil {
				log.Fatalf("Failed to stop server: %v", err)
			}
			return nil
		},
	})
}
`

var appModuleContent = `package app

import (
	"yourmodule/domain/user"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type AppModule struct {
	FiberApp   *fiber.App
	UserModule *user.UserModule
}

func NewAppModule(
	fiberApp *fiber.App,
	userModule *user.UserModule,
) *AppModule {
	return &AppModule{
		FiberApp:   fiberApp,
		UserModule: userModule,
	}
}

var Module = fx.Options(
	fx.Provide(NewAppModule),
	fx.Provide(func() *fiber.App {
		return fiber.New()
	}),
	user.Module,
)
`

var usersControllerContent = `package user

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userService *UserService
}

func NewUserController(userService *UserService) *UserController {
	return &UserController{userService: userService}
}

func (uc *UserController) GetAllUsers(c *fiber.Ctx) error {
	return uc.userService.GetAllUsers(c)
}

func (uc *UserController) GetUser(c *fiber.Ctx) error {
	return uc.userService.GetUser(c)
}

// RegisterRoutes se encarga de registrar las rutas en el módulo.
func RegisterRoutes(app *fiber.App, controller *UserController) {
	fmt.Println("Registering routes...")
	app.Get("/users", controller.GetAllUsers)
	app.Get("/users/:id", controller.GetUser)
}
`

var usersServiceContent = `package user

import (
	"github.com/gofiber/fiber/v2"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) GetAllUsers(c *fiber.Ctx) error {
	users := []string{"User1", "User2", "User3"}
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Fetched all users",
		"data":    users,
	})
}

func (s *UserService) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Fetched user by ID",
		"id":      id,
	})
}
`

var usersModuleContent = `package user

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type UserModule struct {
	UserController *UserController
	UserService    *UserService
}

func NewUserModule(
	app *fiber.App,
	userController *UserController,
	userService *UserService,
) *UserModule {
	module := &UserModule{
		UserController: userController,
		UserService:    userService,
	}

	return module
}

var Module = fx.Options(
	fx.Provide(NewUserModule, NewUserController, NewUserService),
	fx.Invoke(RegisterRoutes),
)
`
