!SLIDE
# Security

!SLIDE
# match do not catch all
    @@@ ruby
    # Rails 3
    match 'users/new' => 'users#new'

    # Rails 4
    match 'users/new' => 'users#new', via: [:get]
    get 'users/new' => 'users#new'

!SLIDE
* html entities escaped by default
* new security headers
* Google security changes

!SLIDE bullets
# Minor changes #

* New Default Test Locations test/models, test/controllers
* Your app's executables now live in the bin
* multiple routes file (#draw method)
* Thread safe on by default
* PATCH verb support
* Object#try will now return nil instead of raise a NoMethodError if the receiving object does not implement the method, but you can still get the old behavior by using the new Object#try!.
* Model.all now returns an ActiveRecord::Relation, rather than an array of records.
* All dynamic methods except for find_by_... and find_by_...! are deprecated. Here's how you can rewrite the code:

!SLIDE bullets
# PostgreSQL #

!SLIDE bullets
# Feature #
