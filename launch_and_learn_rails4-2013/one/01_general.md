!SLIDE
# Launch and learn #
## Rails 4 ##

!SLIDE bullets incremental
# Assumptions #

* Ruby 2.0 preferred; 1.9.3+ required
* less revolutionary then 2.x => 3.0
* New depracation policy

!SLIDE
# General #

!SLIDE
# Strong parameters #

!SLIDE
## Rails 3 way 1/2##

    @@@ ruby
    # app/models/user.rb
    class User
      attr_accessible :name
    end

    # app/controllers/users_controller.rb
    def create
      @user.create(params[:user])
    end

!SLIDE
## Rails 3 way 2/2 ##
    @@@ ruby
    # app/models/user.rb
    class User
      attr_accessible :name, :is_admin, :as => :admin
    end

    # app/controllers/admin/users_controller.rb
    def create
      @user.create(params[:user], :as => :admin)
    end

!SLIDE
## Rails 4 way 1/2##
    @@@ruby
    # app/controllers/users_controller.rb
    def create
      @user.create(user_params)
    end

    private

    # Never trust parameters from the scary internet,
    # only allow the white list through.
    def user_params
      params.require(:user).permit(:name)
    end

!SLIDE
## Rails 4 way 2/2##
    @@@ruby
    # app/controllers/admin/users_controller.rb
    def create
      @user.create(user_params)
    end

    private

    # Never trust parameters from the scary internet,
    # only allow the white list through.
    def user_params
      params.require(:user).permit(:name, :is_admin)
    end

!SLIDE
## Routing concerns 1/2##

    @@@ ruby
    concern :commentable do
      resources :comments
    end

    resources :messages, concerns: :commentable
    resources :posts, concerns: :commentable

!SLIDE
## Routing concerns 2/2 ##

    @@@ ruby
    resources :messages do
      resources :comments
    end

    resources :posts do
      resources :comments
    end

!SLIDE
# Live Streaming
## ActionController::Live
### Server-Sent Events (SSE)

    @@@ javascript
    id: 12345\n
    event: some_channel\n
    data: {"hello":"world"}\n\n

!SLIDE
# Live demo!

!SLIDE
# Russian doll caching

!SLIDE
* declarative ETags
* ActiveModel::Model
* Queue API
* Async Mailer
* Dalli
* Rails::Plugin has gone. Instead of adding plugins to vendor/plugins use gems or bundler with path or git dependencies.
* Turbolinks
