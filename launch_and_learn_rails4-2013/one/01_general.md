!SLIDE
# Launch and learn #
## Rails 4 ##

!SLIDE bullets incremental
# Assumptions #
  * Ruby 2.0 preferred; 1.9.3+ required
  * less revolutionary then 2.x => 3.0

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

!SLIDE bullets incremental smaller
# ActiveModel::Model
* a module that allows Ruby objects to work with Action Pack
* validations
* callbacks
* serialization
* internationalization (i18n) support
* form_for

!SLIDE
    @@@ ruby
    class Contact
      include ActiveModel::Model

      attr_accessor :name, :email, :message

      validates :name, presence: true
      validates :email, presence: true
      validates :message, presence: true, length: { maximum: 300 }
    end

!SLIDE
    @@@ html
    <h1>Contact Us</h1>

    <%= form_for @contact do |f| %>
      <%= f.label :name %>
      <%= f.text_field :name %>

      <%= f.label :email %>
      <%= f.email_field :email %>

      <%= f.label :message %>
      <%= f.text_area :message %>

      <%= f.submit 'Submit' %>
    <% end %>

!SLIDE
    @@@ ruby
    class ContactsController < ApplicationController
      def new
        @contact = Contact.new
      end

      def create
        @contact = Contact.new(params[:contact])
        if @contact.valid?
          ContactMailer.new_contact(@contact).deliver
          redirect_to root_path
        else
          render :new
        end
      end
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
# Live Streaming #
## ActionController::Live ##
### Server-Sent Events (SSE) ###

    @@@ javascript
    id: 12345\n
    event: some_channel\n
    data: {"hello":"world"}\n\n

!SLIDE
# Live demo! #

!SLIDE bullets incremental
# Russian doll caching #

  * fancy name for nesting fragment caches
  * don't explicitly invalidate a cache key, just use a new one

!SLIDE
    @@@ ruby
    class Ticket < ActiveRecord::Base
      has_many :comments
    end

    class Comment < ActiveRecord::Base
      belongs_to :ticket, touch: true
    end

!SLIDE
    @@@ html
    <!-- app/views/tickets/show.html.erb -->
    <% cache @ticket do %>
      Ticket: <%= @ticket.name %>
      <%= render @ticket.comments %>
    <% end %>

    <!-- app/views/comments/_comment.html.erb -->
    <% cache comment do %>
      <%= comment.name %>
    <% end %>

!SLIDE
# Cache digest

     @@@ Shell
     views/comments/1-20131220141922
     views/comments/2-20131220141922
     views/tickets/1-20131220141922

!SLIDE
# Cache digest

     @@@ Shell
     views/comments/1-20131220141922/74865fcb3e2752a0928fa4f89b3e4426
     views/comments/2-20131220141922/74865fcb3e2752a0928fa4f89b3e4426
     views/tickets/1-20131220141922/4277f85c137009873c093088ef609e60

!SLIDE bullets incremental
# Turn into gems

* Page caching - actionpack-page_caching
* Action caching - actionpack-action_caching

!SLIDE bullets incremental smaller
# Rails::Plugin has gone
* Instead of adding plugins to vendor/plugins use gems
