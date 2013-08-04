!SLIDE bullets incremental
# Turn into gems #

* activerecord-deprecated_finders (still included in rails 4.0)
* Page and action caching
* ActiveRecord Observers
* ActiveRecord::SessionStore
* url_for :controller / :action
* ActiveResource

!SLIDE bullets
# Security #

* match do not catch all
* html entities escaped by default
* new security headers
* Google security changes

!SLIDE bullets
# Minor changes #

* test/models, test/controllers
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

!SLIDE bullets
# Links #

* http://guides.rubyonrails.org/4_0_release_notes.html
* http://goo.gl/Ta6vQ
* http://tenderlovemaking.com/2012/07/30/is-it-live.html
