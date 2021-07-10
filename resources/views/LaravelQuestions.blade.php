@extends('layouts.main')

@section('title', 'Laravel Questions')

@section('content')


<section id="about" class="about d-flex align-items-center">
    
    <div class="container">

<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;"><span lang="en-US"><strong>Points to be remembered while appearing in Laravel Interviews :</strong></span></p>
<ul>
    <li>
        <p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>Laravel</strong><span lang="en-US">&nbsp;is developed on the MVC (Model-View-Controller) design pattern.</span></p>
    </li>
    <li>
        <p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;"><strong>Laravel&nbsp;</strong><span lang="en-US">comes with inbuilt features/ modules like authentication, authorization, localization, models, views, sessions, paginations and routing</span></p>
    </li>
    <li>
        <p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;"><strong>Laravel&nbsp;</strong><span lang="en-US">supports advanced concepts of PHP and OOPs like Dependency Injection, traits, Contracts, bundles, Namespaces, Facades</span></p>
    </li>
    <li>
        <p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;"><strong>Laravel&nbsp;</strong>supports Multiple Databases like MySQL,&nbsp;<span lang="en-US"><u>PostgreSQL</u></span><span lang="en-US">, SQLite, SQL Server.</span></p>
    </li>
    <li>
        <p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;"><strong>Laravel&nbsp;</strong><span lang="en-US">allows developers to write clean and modular code.</span></p>
    </li>
    <li>
        <p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;"><strong>Laravel&nbsp;</strong><span lang="en-US">supports blade Template Engine</span></p>
    </li>
    <li>
        <p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;"><strong>Laravel&nbsp;</strong><span lang="en-US">comes with Official Packages like Cashier, Envoy, Horizon, Passport, Scout, Socialite</span></p>
    </li>
    <li>
        <p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;"><strong>Laravel&nbsp;</strong>can be used with various popular Javascript Frameworks Like&nbsp;<span lang="en-US"><u>AngularJs</u></span><span lang="en-US">, VueJs, ReactJS.</span></p>
    </li>
</ul>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><span lang="en-US"><strong>Frequently Asked Laravel Developer Interview Questions.</strong></span></p>
<ol>
    <li>
        <p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><span lang="en-US">What is Laravel&nbsp;migration?</span></p>
    </li>
</ol>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-left: 1.27cm;margin-top: 0.49cm;"><strong>Migration in Laravel</strong> is used to modify and share the structure of your database. It is paired with the schema builder in Laravel to build the database for your application. It acts as a version control but for the Laravel database. There are two methods in the migration class, up and down. Up method is used to add tables, columns, to the database. The down method is used to reverse the operations of up. To create a migration, use&nbsp;<strong>make:migration</strong><span lang="en-US">&nbsp;command. To run the migration, use migrate command. It also offers a rollback feature.</span></p>
<ol start="2">
    <li>
        <p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><span lang="en-US">Which version of Laravel you have used?</span></p>
    </li>
</ol>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-left: 1.27cm;margin-top: 0.49cm;"><span lang="en-US">Answer the versions of Laravel you have worked on like Laravel 4, Laravel 5, Laravel 5.5, Laravel 6</span></p>
<ol start="3">
    <li>
        <p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><span lang="en-US">What is socialite Laravel?</span></p>
    </li>
</ol>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-left: 1.27cm;margin-top: 0.49cm;"><strong>Laravel Socialite</strong><span lang="en-US">&nbsp;is an OAuth provider that is used to authenticate social media apps such as Facebook, Twitter, Google, and many more simply and seamlessly. It can be added to your Laravel web application using the dependency manager composer.</span></p>
<ol start="4">
    <li>
        <p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><span lang="en-US">Why the composer is used in Laravel?</span></p>
    </li>
</ol>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-left: 1.27cm;margin-top: 0.49cm;">The&nbsp;<strong>composer in Laravel</strong><span lang="en-US">&nbsp;is a dependency management tool that is used to install and manage libraries in your application. It does not install packages globally but on a per-application basis. Define the packages you want in the composer.json file and run the following command to install the packages in your application.</span></p>
<ol start="5">
    <li>
        <p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><span lang="en-US">How can you update Laravel?</span></p>
    </li>
</ol>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-left: 1.27cm;margin-top: 0.49cm;"><span lang="en-US">To upgrade your Laravel framework to the latest version, open the composer.json file and change the version of the Laravel framework to the newest version. Then run the following command to update your Laravel framework.</span></p>
<p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;margin-left: 1.27cm;"><span lang="en-US">composer update</span></p>
<ol start="6">
    <li>
        <p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><span lang="en-US">How to use Laravel&nbsp;tinker?</span></p>
    </li>
</ol>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-left: 1.27cm;margin-top: 0.49cm;"><strong>Laravel Tinker</strong><span lang="en-US">&nbsp;is a REPL that allows the developer to interact with the Laravel application through the command line. To enter into the Tinker environment, enter the following command</span></p>
<p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;margin-left: 1.27cm;"><span lang="en-US">php artisan tinker</span></p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-left: 1.27cm;margin-top: 0.49cm;"><span lang="en-US">Tinker allows you to run commands such as clear-compiled, down, env, inspire, migrate, optimize, and up by default. To use more commands using Tinker, add the command to the command array that is present in the tinker.php configuration file.</span></p>
<ol start="7">
    <li>
        <p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><span lang="en-US">How to use where relationship Laravel?</span></p>
    </li>
</ol>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-left: 1.27cm;margin-top: 0.49cm;">The&nbsp;<strong>where clause</strong><span lang="en-US">&nbsp;is a query builder that is used to verify something. Where has three arguments that it must require. The column name is the first argument with an operator as the second argument. The third argument is the value that the first operand must be evaluated using the operator.</span></p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-left: 1.27cm;margin-top: 0.49cm;"><span lang="en-US"><strong>Example</strong></span></p>
<p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;margin-left: 1.27cm;">$users = DB::table(&apos;users&apos;)-&gt;where(&apos;salary&apos;, &apos;=&apos;, 100)-&gt;get();&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-left: 1.27cm;margin-top: 0.49cm;"><span lang="en-US">In the above example, the where clause is used to get the users who have the salary column that is equal to 100.</span></p>
<ol start="8">
    <li>
        <p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><span lang="en-US">What is the difference between Laravel&nbsp;find and where?</span></p>
    </li>
</ol>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-left: 1.27cm;margin-top: 0.49cm;"><span lang="en-US">The where clause in the query is used to verify and retrieve values. It has three arguments. First is the column name argument. The second is the operator. The third is the values that need to be checked with the column name. It retrieves the first model that is matching the given query.</span></p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-left: 1.27cm;margin-top: 0.49cm;"><span lang="en-US"><strong>Example</strong></span></p>
<p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;margin-left: 1.27cm;">$users = DB::table(&apos;students&apos;)-&gt;where(&apos;marks&apos;, &apos;=&apos;, 100)-&gt;get();&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-left: 1.27cm;margin-top: 0.49cm;"><span lang="en-US">The find method is used to retrieve a single model instance instead of a collection of models. It retrieves the model by its primary key.</span></p>
<p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;margin-left: 1.27cm;"><span lang="en-US">$users = StudentModel::find(1);</span></p>
<ol start="9">
    <li>
        <p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><span lang="en-US">How to hash password in Laravel?</span></p>
    </li>
</ol>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-left: 1.27cm;margin-top: 0.49cm;"><span lang="en-US">The Hash::make function is used to create a hash for the password.</span></p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-left: 1.27cm;margin-top: 0.49cm;"><span lang="en-US"><strong>Syntax:</strong></span></p>
<p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;margin-left: 1.27cm;"><span lang="en-US">Hash::make($password)</span></p>
<ol start="10">
    <li>
        <p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><span lang="en-US">What are the Laravel guards?</span></p>
    </li>
</ol>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-left: 1.27cm;margin-top: 0.49cm;">The&nbsp;<strong>guard in Laravel</strong><span lang="en-US">&nbsp;is used to define the user authentication for each request. The Guards also provide a definition of user information storage and retrieval by the system. The session guard is a default guard in Laravel that is used to maintain the state using session storage and cookies. A web guard is used to store and retrieve session information. API guard is used for authenticating users and requests. These are some examples of default Laravel guard, and you can also create your own guard in Laravel.</span></p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><br>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>1) What is Laravel?&nbsp;</strong></p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">Laravel is an open-source widely used PHP framework. The platform was intended for the development of web application by using MVC architectural pattern. Laravel is released under the MIT license.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">Therefore, its source code is hosted on GitHub. It is a reliable PHP framework as it follows expressive and accurate language rules.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>2) Define composer.</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">It is an application-level package manager for PHP. It provides a standard format for managing PHP software dependencies and libraries.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>3) What is HTTP middleware?</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">HTTP middleware is a technique for filtering HTTP requests. Laravel includes a middleware that checks whether application user is authenticated or not.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>4) Name aggregates methods of query builder.</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><span lang="en-US">Aggregates methods of query builder are: 1) max(), 2) min(), 3) sum(), 4) avg(), and 5) count().</span></p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>5) What is a Route?</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">A route is basically an endpoint specified by a URI (Uniform Resource Identifier). It acts as a pointer in Laravel application.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">Most commonly, a route simply points to a method on a controller and also dictates which HTTP methods are able to hit that URI.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>6) Why use Route?</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">Routes are stored inside files under the /routes folder inside the project&apos;s root directory. By default, there are a few different files corresponding to the different &quot;sides&quot; of the application (&quot;sides&quot; comes from the hexagonal architecture methodology).&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>7) What do you mean by bundles?</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">In Laravel, bundles are referred to as packages. These packages are used to increase the functionality of Laravel. A package can have views, configuration, migrations, routes, and tasks.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>8) Explain important directories used in a common Laravel application.</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">Directories used in a common Laravel application are:&nbsp;</p>
<ul>
    <li>
        <p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><span lang="en-US">App/: This is a source folder where our application code lives. All controllers, policies, and models are inside this folder.</span></p>
    </li>
    <li>
        <p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;"><span lang="en-US">Config/: Holds the app&apos;s configuration files. These are usually not modified directly but instead, rely on the values set up in the .env (environment) file at the root of the app.</span></p>
    </li>
    <li>
        <p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;"><span lang="en-US">Database/: Houses the database files, including migrations, seeds, and test factories.</span></p>
    </li>
    <li>
        <p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;"><span lang="en-US">Public/: Publicly accessible folder holding compiled assets and of course an index.php file.</span></p>
    </li>
</ul>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>9) What is a Controller?</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">A controller is the &quot;C&quot; in the &quot;MVC&quot; (Model-View-Controller) architecture, which is what Laravel is based on.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>10) Explain reverse routing in Laravel.</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">Reverse routing is a method of generating URL based on symbol or name. It makes your Laravel application flexible.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>12) Explain traits in Laravel.</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">Laravel traits are a group of functions that you include within another class. A trait is like an abstract class. You cannot instantiate directly, but its methods can be used in concreate class.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>13) Explain the concept of contracts in Laravel.</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">They are set of interfaces of Laravel framework. These contracts provide core services. Contracts defined in Laravel include corresponding implementation of framework.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>14) How will you register service providers?&nbsp;</strong></p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">You can register service providers in the config/app.php configuration file that contains an array where you can mention the service provider class name.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>15) Where will you define Laravel&apos;s Facades?</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">All facades of Laravel have defined in Illuminate\Support\Facades namespace.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>16) State the difference between get and post method.</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">Get method allows you to send a limited amount of data in the header. Post allows you to send a large amount of data in the body.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><span lang="en-US"><strong>17) List default packages of Laravel 5.6.</strong></span></p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">Default packages of Laravel 5.6 are: 1) Envoy, 2) Passport, 3) Socialite, 4) Cashier, 5) Horizon, and 6) Scout.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>18) What is service container in Laravel?</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">Service container is a tool used for performing dependency injection in Laravel.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>19) How can you enable query log in Laravel?</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">You can use enableQueryLog method to enable query log in Laravel.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>20) Explain the concept of events in Laravel.</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">An event is an occurrence or action that help you to subscribe and listen for events that occur in Laravel application. Some of the events are fired automatically by Laravel when any activity occurs.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>21) Explain dependency injection and their types.&nbsp;</strong></p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">It is a technique in which one object is dependent on another object. There are three types of dependency injection: 1) Constructor injection, 2) setter injection, and 3) interface injection.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>22) What are the advantages of using Laravel?</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">Here are important benefits of Laravel:&nbsp;</p>
<ul>
    <li>
        <p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><span lang="en-US">Laravel has blade template engine to create dynamic layouts and increase compiling tasks.</span></p>
    </li>
    <li>
        <p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;"><span lang="en-US">Reuse code without any hassle.</span></p>
    </li>
    <li>
        <p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;"><span lang="en-US">Laravel provides you to enforce constraints between multiple DBM objects by using an advanced query builder mechanism.</span></p>
    </li>
    <li>
        <p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;"><span lang="en-US">The framework has an auto-loading feature, so you don&apos;t do manual maintenance and inclusion paths</span></p>
    </li>
    <li>
        <p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;"><span lang="en-US">The framework helps you to make new tools by using LOC container.</span></p>
    </li>
    <li>
        <p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;"><span lang="en-US">Laravel offers a version control system that helps with simplified management of migrations.</span></p>
    </li>
</ul>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>23) Explain validation concept in Laravel.</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">Validations are an important concept while designing any Laravel application. It ensures that the data is always in an expected format before it stores into the database. Laravel provides many ways to validate your data.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">Base controller trait uses a ValidatesRequests class which provides a useful method to validate requests coming from the client machine.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>24) What does ORM stand for?</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">ORM stands for Object Relational Mapping&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>25) How can you reduce memory usage in Laravel?</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">While processing a large amount of data, you can use the cursor method in order to reduce memory usage.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><span lang="en-US"><strong>26) List available types of relationships in Laravel Eloquent.</strong></span></p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">Types of relationship in Laravel Eloquent are: 1) One To One 2) One To Many 3) Many To Many 4) Has Many Through, and 5) Polymorphic Relations.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>27) Name the Template Engine utilized by Laravel.</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">Blade is a powerful template engine utilized by Laravel.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>28) Name databases supported by Laravel.</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">Laravel supports the following databases:&nbsp;</p>
<ul>
    <li>
        <p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><span lang="en-US">PostgreSQL</span></p>
    </li>
    <li>
        <p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;"><span lang="en-US">SQL Server</span></p>
    </li>
    <li>
        <p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;"><span lang="en-US">SQLite</span></p>
    </li>
    <li>
        <p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;"><span lang="en-US">MySQL</span></p>
    </li>
</ul>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>29) Why are migrations important?</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">Migrations are important because it allows you to share application by maintaining database consistency. Without migration, it is difficult to share any Laravel application. It also allows you to sync database.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>30) Define Lumen</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">Lumen is a micro-framework. It is a smaller, and faster, version of a building Laravel based services, and REST API&apos;s.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>31) Explain PHP artisan</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">An artisan is a command-line tool of Laravel. It provides commands that help you to build Laravel application without any hassle.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>32) How can you generate URLs?</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">Laravel has helpers to generate URLs. This is helpful when you build link in your templates and API response.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>33) Which class is used to handle exceptions?</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">Laravel exceptions are handled by App\Exceptions\Handler class.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>34) What are common HTTP error codes?</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">The most common HTTP error codes are:&nbsp;</p>
<ul>
    <li>
        <p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><span lang="en-US">Error 404 &ndash; Displays when Page is not found.</span></p>
    </li>
    <li>
        <p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;"><span lang="en-US">Error- 401 &ndash; Displays when an error is not authorized</span></p>
    </li>
</ul>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>35) Explain fluent query builder in Laravel.</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">It is a database query builder that provides convenient, faster interface to create and run database queries.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>36) What is the use of dd() function?</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">This function is used to dump contents of a variable to the browser. The full form of dd is Dump and Die.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>37) List out common artisan commands used in Laravel.</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">Laravel supports following artisan commands:&nbsp;</p>
<ul>
    <li>
        <p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><span lang="en-US">PHP artisan down;</span></p>
    </li>
    <li>
        <p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;"><span lang="en-US">PHP artisan up;</span></p>
    </li>
    <li>
        <p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;"><span lang="en-US">PHP artisan make:controller;</span></p>
    </li>
    <li>
        <p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;"><span lang="en-US">PHP artisan make:model;</span></p>
    </li>
    <li>
        <p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;"><span lang="en-US">PHP artisan make:migration;</span></p>
    </li>
    <li>
        <p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;"><span lang="en-US">PHP artisan make:middleware;</span></p>
    </li>
</ul>
php artisan down --render="errors::illustrated-layout" <br>
This will copy the error views from /vendor/laravel/framework/src/Illuminate/Foundation/Exceptions/views to your project’s /resources/views/errors. Here are all the error views that you’ll get.

<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>38) How to configure a mail-in Laravel?</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">Laravel provides APIs to send an email on local and live server.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>39) Explain Auth.</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">It is a method of identifying user login credential with a password. In Laravel it can be managed with a session which takes two parameters 1) username and 2) password.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>40) Differentiate between delete() and softDeletes().</strong>&nbsp;</p>
<ul>
    <li>
        <p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><span lang="en-US">delete(): remove all record from the database table.</span></p>
    </li>
    <li>
        <p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;"><span lang="en-US">softDeletes(): It does not remove the data from the table. It is used to flag any record as deleted.</span></p>
    </li>
</ul>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>41) How can you make real time sitemap.xml file in Laravel?</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">You can create all web pages of a website to tell the search engine about the organizing site content. The crawlers of search engine read this file intelligently to crawl a website.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>42) Explain faker in Laravel.</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">It is a type of module or packages which are used to create fake data. This data can be used for testing purpose.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">It is can also be used to generate: 1) Numbers, 2) Addresses, 3) DateTime, 4) Payments, and 5) Lorem text.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>43) How will you check table is exists or in the database?</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">Use hasTable() Laravel function to check the desired table is exists in the database or not.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>44) What is the significant difference between insert() and insertGetId() function in Laravel?</strong>&nbsp;</p>
<ul>
    <li>
        <p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><span lang="en-US">Insert(): This function is simply used to insert a record into the database. It not necessary that ID should be autoincremented.&nbsp;</span></p>
    </li>
    <li>
        <p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;"><span lang="en-US">InsertGetId(): This function also inserts a record into the table, but it is used when the ID field is auto-increment.</span></p>
    </li>
</ul>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>45) Explain active record concept in Laravel.</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">In active record, class map to your database table. It helps you to deal with CRUD operation.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>46) List basic concepts in Laravel?</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">Following are basic concepts used in Laravel:&nbsp;</p>
<ul>
    <li>
        <p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><span lang="en-US">Routing</span></p>
    </li>
    <li>
        <p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;"><span lang="en-US">Eloquent ORM</span></p>
    </li>
    <li>
        <p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;"><span lang="en-US">Middleware</span></p>
    </li>
    <li>
        <p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;"><span lang="en-US">Security</span></p>
    </li>
    <li>
        <p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;"><span lang="en-US">Caching</span></p>
    </li>
    <li>
        <p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;"><span lang="en-US">Blade Templating</span></p>
    </li>
</ul>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>47) Define Implicit Controller.</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">Implicit Controllers help you to define a proper route to handle controller action. You can define them in route.php file with Route:: controller() method.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>48) How to use the custom table in Laravel Model?</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">In order to use a custom table, you can override the property of the protected variable $table.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>49) What is MVC framework?</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">It is Model, View, and Controller:&nbsp;</p>
<ul>
    <li>
        <p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><span lang="en-US">Model: Model defines logic to write Laravel application.</span></p>
    </li>
    <li>
        <p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;"><span lang="en-US">View: It covers UI logic of Laravel application.</span></p>
    </li>
    <li>
        <p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;"><span lang="en-US">Controller: It is work as an interface between Model, and View. It is a way how the user interacts with an application.</span></p>
    </li>
</ul>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>50) Define '@@include'.</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">@@include is used to load more than one template view files. It helps you to include view within another view. User can also load multiple files in one view.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>51) Explain the concept of cookies.</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">Cookies are small file sent from a particular website and stored on PC by user&apos;s browser while the user is browsing.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>52) Which file is used to create a connection with the database?</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">To create a connection with the database, you can use .env file.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>53) What is Eloquent?</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">Eloquent is an ORM used in Laravel. It provides simple active record implementation working with the database. Each database table has its Model, which used to interact with the table.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>54) Name some Inbuilt Authentication Controllers of Laravel.</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">Laravel installation has an inbuilt set of common authentication controllers. These controllers are:&nbsp;</p>
<ul>
    <li>
        <p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><span lang="en-US">RegisterController</span></p>
    </li>
    <li>
        <p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;"><span lang="en-US">LoginController</span></p>
    </li>
    <li>
        <p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;"><span lang="en-US">ResetPasswordController</span></p>
    </li>
    <li>
        <p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;"><span lang="en-US">ForgetPasswordController</span></p>
    </li>
</ul>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>55) Define Laravel guard.</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">Laravel guard is a special component that is used to find authenticated users. The incoming requested is initially routed through this guard to validate credentials entered by users. Guards are defined in ../config/auth.php file.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>56) What is Laravel API rate limit?</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">It is a feature of Laravel. It provides handle throttling. Rate limiting helps Laravel developers to develop a secure application and prevent DOS attacks.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>57) Explain collections in Laravel.</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">Collections is a wrapper class to work with arrays. Laravel Eloquent queries use a set of the most common functions to return database result.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>58) What is the use of DB facade?</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">DB facade is used to run SQL queries like create, select, update, insert, and delete.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>59) What is the use of Object Relational Mapping?</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">Object Relational Mapping is a technique that helps developers to address, access, and manipulate objects without considering the relation between object and their data sources.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>60) Explain the concept of routing in Laravel.</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">It allows routing all your application requests to the controller. Laravel routing acknowledges and accepts a Uniform Resource Identifier with a closure.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>61) What is Ajax in Laravel?</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">Ajax stands for Asynchronous JavaScript and XML is a web development technique that is used to create asynchronous Web applications. In Laravel, response() and json() functions are used to create asynchronous web applications.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>62) What is a session in Laravel?</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">Session is used to pass user information from one web page to another. Laravel provides various drivers like a cookie, array, file, Memcached, and Redis to handle session data.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>63) How to access session data?</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">Session data be access by creating an instance of the session in HTTP request. Once you get the instance, use get() method with a &quot;Key&quot; as a parameter to get the session details.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>64) State the difference between authentication and authorization.</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">Authentication means confirming user identities through credentials, while authorization refers to gathering access to the system.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>65) Explain to listeners.</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">Listeners are used to handling events and exceptions. The most common listener in Laravel for login event is LoginListener.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>66) What are policies classes?</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">Policies classes include authorization logic of Laravel application. These classes are used for a particular model or resource.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>67) How to rollback last migration?</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">Use need to use artisan command to rollback the last migration.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>68) What do you mean by Laravel Dusk?</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">Laravel Dusk is a tool which is used for testing JavaScript enabled applications. It provides powerful, browser automation, and testing API.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>69) Explain Laravel echo.</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">It is a JavaScript library that makes possible to subscribe and listen to channels Laravel events. You can use NPM package manager to install echo.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>70) What is make method?</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">Laravel developers can use make method to bind an interface to concreate class. This method returns an instance of the class or interface. Laravel automatically inject dependencies defined in class constructor.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>71) Explain Response in Laravel.</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">All controllers and routes should return a response to be sent back to the web browser. Laravel provides various ways to return this response. The most basic response is returning a string from controller or route.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>72) What is query scope?</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">It is a feature of Laravel where we can reuse similar queries. We do not require to write the same types of queries again in the Laravel project. Once the scope is defined, just call the scope method when querying the model.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>73) Explain homestead in Laravel.</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">Laravel homestead is the official, disposable, and pre-packaged vagrant box that a powerful development environment without installing HHVM, a web server, and PHP on your computer.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>74) What is namespace in Laravel?</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">A namespace allows a user to group the functions, classes, and constants under a specific name.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>75) What is Laravel Forge?</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">Laravel Forge helps in organizing and designing a web application. Although the manufacturers of the Laravel framework developed this toll, it can automate the deployment of every web application that works on a PHP server.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>76) State the difference between CodeIgniter and Laravel.</strong>&nbsp;</p>
<table cellpadding="1" cellspacing="0" width="611">
    <tbody>
        <tr>
            <td style="border: none;padding: 0cm;" width="21.818181818181817%">
                <p style="margin-bottom: 0.25cm;line-height: 115%;text-align: left;background: transparent;"><strong>Parameter</strong></p>
            </td>
            <td style="border: none;padding: 0cm;" width="36.85950413223141%">
                <p style="margin-bottom: 0.25cm;line-height: 115%;text-align: left;background: transparent;"><strong>CodeIgniter</strong></p>
            </td>
            <td style="border: none;padding: 0cm;" width="41.32231404958678%">
                <p style="margin-bottom: 0.25cm;line-height: 115%;text-align: left;background: transparent;"><strong>Laravel</strong></p>
            </td>
        </tr>
        <tr>
            <td style="border: none;padding: 0cm;" width="21.818181818181817%">
                <p style="margin-bottom: 0.25cm;line-height: 115%;text-align: left;background: transparent;"><span lang="en-US">Support of ORM&nbsp;</span></p>
            </td>
            <td style="border: none;padding: 0cm;" width="36.85950413223141%">
                <p style="margin-bottom: 0.25cm;line-height: 115%;text-align: left;background: transparent;"><span lang="en-US">CodeIgniter does not support Object-relational mapping.&nbsp;</span></p>
            </td>
            <td style="border: none;padding: 0cm;" width="41.32231404958678%">
                <p style="margin-bottom: 0.25cm;line-height: 115%;text-align: left;background: transparent;"><span lang="en-US">Laravel supports ORM.&nbsp;</span></p>
            </td>
        </tr>
        <tr>
            <td style="border: none;padding: 0cm;" width="21.818181818181817%">
                <p style="margin-bottom: 0.25cm;line-height: 115%;text-align: left;background: transparent;"><span lang="en-US">Provide Authentication&nbsp;</span></p>
            </td>
            <td style="border: none;padding: 0cm;" width="36.85950413223141%">
                <p style="margin-bottom: 0.25cm;line-height: 115%;text-align: left;background: transparent;"><span lang="en-US">It does provide user authentication.&nbsp;</span></p>
            </td>
            <td style="border: none;padding: 0cm;" width="41.32231404958678%">
                <p style="margin-bottom: 0.25cm;line-height: 115%;text-align: left;background: transparent;"><span lang="en-US">It has inbuilt user authentication.&nbsp;</span></p>
            </td>
        </tr>
        <tr>
            <td style="border: none;padding: 0cm;" width="21.818181818181817%">
                <p style="margin-bottom: 0.25cm;line-height: 115%;text-align: left;background: transparent;"><span lang="en-US">Programming Paradigm&nbsp;</span></p>
            </td>
            <td style="border: none;padding: 0cm;" width="36.85950413223141%">
                <p style="margin-bottom: 0.25cm;line-height: 115%;text-align: left;background: transparent;"><span lang="en-US">It is component-oriented.&nbsp;</span></p>
            </td>
            <td style="border: none;padding: 0cm;" width="41.32231404958678%">
                <p style="margin-bottom: 0.25cm;line-height: 115%;text-align: left;background: transparent;"><span lang="en-US">It is object-oriented.&nbsp;</span></p>
            </td>
        </tr>
        <tr>
            <td style="border: none;padding: 0cm;" width="21.818181818181817%">
                <p style="margin-bottom: 0.25cm;line-height: 115%;text-align: left;background: transparent;"><span lang="en-US">Support of other Database Management System&nbsp;</span></p>
            </td>
            <td style="border: none;padding: 0cm;" width="36.85950413223141%">
                <p style="margin-bottom: 0.25cm;line-height: 115%;text-align: left;background: transparent;"><span lang="en-US">It supports Microsoft SQL Server, ORACLE, MYSQL, IBM DB2, PostgreSQL, JDBC, and orientDB compatible.&nbsp;</span></p>
            </td>
            <td style="border: none;padding: 0cm;" width="41.32231404958678%">
                <p style="margin-bottom: 0.25cm;line-height: 115%;text-align: left;background: transparent;"><span lang="en-US">It supports PostgreSQL, MySQL, MongoDB, and Microsoft BI, but CodeIgniter additionally supports other databases like Microsoft SQL Server, DB2, Oracle, etc.&nbsp;</span></p>
            </td>
        </tr>
        <tr>
            <td style="border: none;padding: 0cm;" width="21.818181818181817%">
                <p style="margin-bottom: 0.25cm;line-height: 115%;text-align: left;background: transparent;"><span lang="en-US">HTTPS Support&nbsp;</span></p>
            </td>
            <td style="border: none;padding: 0cm;" width="36.85950413223141%">
                <p style="margin-bottom: 0.25cm;line-height: 115%;text-align: left;background: transparent;"><span lang="en-US">CodeIgniter partially support HTTPS. Therefore, programmers can use the URL to secure the data transmission process by creating PATS.&nbsp;</span></p>
            </td>
            <td style="border: none;padding: 0cm;" width="41.32231404958678%">
                <p style="margin-bottom: 0.25cm;line-height: 115%;text-align: left;background: transparent;"><span lang="en-US">Laravel supports custom HTTPS routes. The programmers can create a specific URL for HTTPS route they have defined.&nbsp;</span></p>
            </td>
        </tr>
    </tbody>
</table>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>77) What is an Observer?</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">Model Observers is a feature of Laravel. It is used to make clusters of event listeners for a model. Method names of these classes depict the Eloquent event. Observers classes methods receive the model as an argument.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>78) What is the use of the bootstrap directory?</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">It is used to initialize a Laravel project. This bootstrap directory contains app.php file that is responsible for bootstrapping the framework.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>79) What is the default session timeout duration?</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">The default Laravel session timeout duration is 2 hours.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>80) How to remove a complied class file?</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">Use clear-compiled command to remove the compiled class file.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>81) In which folder robot.txt is placed?</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">Robot.txt file is placed in Public directory.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>82) Explain API.PHP route.</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">Its routes correspond to an API cluster. It has API middleware which is enabled by default in Laravel. These routes do not have any state and cross-request memory or have no sessions.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>83) What is named route?</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">Name route is a method generating routing path. The chaining of these routes can be selected by applying the name method onto the description of route.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>84) what is open source software?</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">Open-source software is a software which source code is freely available. The source code can be shared and modified according to the user requirement.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>85) Explain Loggin in Laravel.</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">It is a technique in which system log generated errors. Loggin is helpful to increase the reliability of the system. Laravel supports various logging modes like syslog, daily, single, and error log modes.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>86) What is Localization?</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">It is a feature of Laravel that supports various language to be used in the application. A developer can store strings of different languages in a file, and these files are stored at resources/views folder. Developers should create a separate folder for each supported language.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>87) Define hashing in Laravel.</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">It is the method of converting text into a key that shows the original text. Laravel uses the Hash facade to store the password securely in a hashed manner.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>88) Explain the concept of encryption and decryption in Laravel.</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">It is a process of transforming any message using some algorithms in such way that the third user cannot read information. Encryption is quite helpful to protect your sensitive information from an intruder.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">Encryption is performed using a Cryptography process. The message which is to be encrypted called as a plain message. The message obtained after the encryption is referred to as cipher message. When you convert cipher text to plain text or message, this process is called as decryption.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>89) How to share data with views?</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">To pass data to all views in Laravel use method called share(). This method takes two arguments, key, and value.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">Generally, share() method are called from boot method of Laravel application service provider. A developer can use any service provider, AppServiceProvider, or our own service provider.&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>90) Explain web.php route.</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">Web.php is the public-facing &quot;browser&quot; based route. This route is the most common and is what gets hit by the web browser. They run through the web middleware group and also contain facilities for CSRF protection (which helps defend against form-based malicious attacks and hacks) and generally contain a degree of &quot;state&quot; (by this I mean they utilize sessions).&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong>91) How to generate a request in Laravel?</strong>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">Use the following artisan command in Laravel to generate request:&nbsp;</p>
<p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;"><span lang="en-US">php artisan make:request UploadFileRequest</span></p>
<p style="margin-bottom: 0.35cm;line-height: 115%;text-align: left;background: transparent;"><br>&nbsp;</p>
<p style="margin-bottom: 0.35cm;line-height: 115%;text-align: left;background: transparent;"><br>&nbsp;</p>
<p style="margin-bottom: 0.35cm;line-height: 115%;text-align: left;background: transparent;"><br>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><span lang="en-US"><strong>What is Laravel?</strong></span></p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">Laravel is a Symfony-based free open-source PHP web framework. It is created by Taylor Otwell and allows developers to write expressive, elegant syntax. Laravel comes with built-in support for user authentication and authorization which is missing in some most popular PHP frameworks like CodeIgniter, CakePHP. It one of the most popular PHP framework which is based on MVC architecture. It was written in PHP &amp; initially released in June 2011. the latest version of laravel is Laravel 8, released on 8 Sep 2020. You will find here the best and latest&nbsp;<strong>Laravel Interview Questions and Answers</strong> that help you crack Laravel Interviews. You can also download the&nbsp;<strong>Laravel Interview Questions PDF&nbsp;</strong><span lang="en-US">from&nbsp;the given link between the post.</span></p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><span lang="en-US"><strong>Quick Questions About Laravel</strong></span></p>
<table cellpadding="1" cellspacing="0" width="611">
    <tbody>
        <tr>
            <td style="border: none;padding: 0cm;" width="30.033003300330034%">
                <p style="margin-bottom: 0.25cm;line-height: 115%;text-align: left;background: transparent;"><strong>Laravel&nbsp;</strong><span lang="en-US">was initially released in&nbsp;</span></p>
            </td>
            <td style="border: none;padding: 0cm;" width="69.96699669966996%">
                <p style="margin-bottom: 0.25cm;line-height: 115%;text-align: left;background: transparent;"><span lang="en-US">June 2011 (About 10 Years ago)</span></p>
            </td>
        </tr>
        <tr>
            <td style="border: none;padding: 0cm;" width="30.033003300330034%">
                <p style="margin-bottom: 0.25cm;line-height: 115%;text-align: left;background: transparent;"><strong>Laravel</strong><span lang="en-US">&nbsp;is written In</span></p>
            </td>
            <td style="border: none;padding: 0cm;" width="69.96699669966996%">
                <p style="margin-bottom: 0.25cm;line-height: 115%;text-align: left;background: transparent;"><span lang="en-US">PHP Programming (PHP 7)</span></p>
            </td>
        </tr>
        <tr>
            <td style="border: none;padding: 0cm;" width="30.033003300330034%">
                <p style="margin-bottom: 0.25cm;line-height: 115%;text-align: left;background: transparent;"><strong>Laravel</strong><span lang="en-US">&nbsp;is a</span></p>
            </td>
            <td style="border: none;padding: 0cm;" width="69.96699669966996%">
                <p style="margin-bottom: 0.25cm;line-height: 115%;text-align: left;background: transparent;"><span lang="en-US">PHP Framework for Developing websites and mobile API&apos;s.</span></p>
            </td>
        </tr>
        <tr>
            <td style="border: none;padding: 0cm;" width="30.033003300330034%">
                <p style="margin-bottom: 0.25cm;line-height: 115%;text-align: left;background: transparent;"><strong>Laravel</strong><span lang="en-US">&nbsp;is developed By</span></p>
            </td>
            <td style="border: none;padding: 0cm;" width="69.96699669966996%">
                <p style="margin-bottom: 0.25cm;line-height: 115%;text-align: left;background: transparent;"><span lang="en-US">Taylor Otwell</span></p>
            </td>
        </tr>
        <tr>
            <td style="border: none;padding: 0cm;" width="30.033003300330034%">
                <p style="margin-bottom: 0.25cm;line-height: 115%;text-align: left;background: transparent;"><strong>Laravel</strong><span lang="en-US">&nbsp;is Based on</span></p>
            </td>
            <td style="border: none;padding: 0cm;" width="69.96699669966996%">
                <p style="margin-bottom: 0.25cm;line-height: 115%;text-align: left;background: transparent;"><span lang="en-US">MVC architectural pattern</span></p>
            </td>
        </tr>
        <tr>
            <td style="border: none;padding: 0cm;" width="30.033003300330034%">
                <p style="margin-bottom: 0.25cm;line-height: 115%;text-align: left;background: transparent;"><strong>Laravel</strong><span lang="en-US">&nbsp;Dependencies</span></p>
            </td>
            <td style="border: none;padding: 0cm;" width="69.96699669966996%">
                <p style="margin-bottom: 0.25cm;line-height: 115%;text-align: left;background: transparent;"><span lang="en-US">Composer, OpenSSL, PDO, Mbstring, etc.</span></p>
            </td>
        </tr>
        <tr>
            <td style="border: none;padding: 0cm;" width="30.033003300330034%">
                <p style="margin-bottom: 0.25cm;line-height: 115%;text-align: left;background: transparent;"><strong>Laravel</strong><span lang="en-US">&nbsp;Licence</span></p>
            </td>
            <td style="border: none;padding: 0cm;" width="69.96699669966996%">
                <p style="margin-bottom: 0.25cm;line-height: 115%;text-align: left;background: transparent;"><span lang="en-US">MIT License</span></p>
            </td>
        </tr>
        <tr>
            <td style="border: none;padding: 0cm;" width="30.033003300330034%">
                <p style="margin-bottom: 0.25cm;line-height: 115%;text-align: left;background: transparent;"><strong>Laravel</strong><span lang="en-US">&nbsp;Current Stable release</span></p>
            </td>
            <td style="border: none;padding: 0cm;" width="69.96699669966996%">
                <p style="margin-bottom: 0.25cm;line-height: 115%;text-align: left;background: transparent;"><span lang="en-US">8.0</span></p>
            </td>
        </tr>
    </tbody>
</table>
<p style="margin-bottom: 0.35cm;line-height: 115%;text-align: left;background: transparent;"><br>&nbsp;</p>
<p style="margin-bottom: 0.35cm;line-height: 115%;text-align: left;background: transparent;"><br>&nbsp;</p>
<h3 style='margin-top: 0.35cm;margin-bottom: 0cm;color: #4f81bd;line-height: 115%;text-align: left;background: transparent;font-family: "Cambria", serif;font-weight: bold;'><strong><span lang="en-US"><u>1) What is Laravel?</u></span></strong><u>&nbsp;</u></h3>
<p style="margin-bottom: 0.35cm;line-height: 115%;text-align: left;background: transparent;"><br>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong><span lang="en-US">Laravel</span></strong><span lang="en-US">&nbsp;is a free open source &quot;PHP framework&quot; based on the MVC design pattern.<br>&nbsp;It is created by Taylor Otwell. Laravel provides expressive and elegant syntax that helps in creating a wonderful web application easily and quickly.</span></p>
<p style="margin-bottom: 0.35cm;line-height: 115%;text-align: left;background: transparent;"><br>&nbsp;</p>
<h3 style='margin-top: 0.35cm;margin-bottom: 0cm;color: #4f81bd;line-height: 115%;text-align: left;background: transparent;font-family: "Cambria", serif;font-weight: bold;'><strong><span lang="en-US"><u>2) Explain Events in laravel ?</u></span></strong><u>&nbsp;</u></h3>
<p style="margin-bottom: 0.35cm;line-height: 115%;text-align: left;background: transparent;"><br>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><span lang="en-US">An event is an action or occurrence recognized by a program that may be handled by the program or code. Laravel events provides a simple observer implementation, that allowing you to subscribe and listen for various events/actions that occur in your application.</span></p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><span lang="en-US">All&nbsp;</span><strong><span lang="en-US">Event</span></strong><span lang="en-US">&nbsp;classes are generally stored in the app/Events directory, while their listeners are stored in app/Listeners of your application.</span></p>
<p style="margin-bottom: 0.35cm;line-height: 115%;text-align: left;background: transparent;"><br>&nbsp;</p>
<h3 style='margin-top: 0.35cm;margin-bottom: 0cm;color: #4f81bd;line-height: 115%;text-align: left;background: transparent;font-family: "Cambria", serif;font-weight: bold;'><strong><span lang="en-US"><u>3) Explain validations in laravel?</u></span></strong><u>&nbsp;</u></h3>
<p style="margin-bottom: 0.35cm;line-height: 115%;text-align: left;background: transparent;"><br>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><span lang="en-US">In Programming validations are a handy way to ensure that your data is always in a clean and expected format before it gets into your database.</span></p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><span lang="en-US">Laravel provides several different ways to validate your application incoming data.By default Laravel&rsquo;s base controller class uses a&nbsp;</span><strong><span lang="en-US">ValidatesRequests</span></strong><span lang="en-US">&nbsp;</span><strong><span lang="en-US">trait</span></strong><span lang="en-US">&nbsp;which provides a convenient method to validate all incoming HTTP requests coming from client.You can also validate data in laravel by creating Form Request.</span></p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong><span lang="en-US">Laravel validation Example</span></strong></p>
<p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;"><span lang="en-US">$validatedData = $request-&gt;validate([</span></p>
<p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;"><span lang="en-US">&apos;name&apos; =&gt; &apos;required|max:255&apos;,</span></p>
<p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;"><span lang="en-US">&apos;username&apos; =&gt; &apos;required|alpha_num&apos;,</span></p>
<p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;"><span lang="en-US">&apos;age&apos; =&gt; &apos;required|numeric&apos;,</span></p>
<p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;"><span lang="en-US">]);</span></p>
<p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;"><br></p>
<p style="margin-bottom: 0.35cm;line-height: 115%;text-align: left;background: transparent;"><br>&nbsp;</p>
<h3 style='margin-top: 0.35cm;margin-bottom: 0cm;color: #4f81bd;line-height: 115%;text-align: left;background: transparent;font-family: "Cambria", serif;font-weight: bold;'><strong><span lang="en-US"><u>4) What is the latest version of Laravel?</u></span></strong><u>&nbsp;</u></h3>
<p style="margin-bottom: 0.35cm;line-height: 115%;text-align: left;background: transparent;"><br>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><span lang="en-US">Laravel 8 is the latest version of Laravel. It was officially released on 11 Sep 2020. Laravel 8 has made with new features like Laravel Jetstream, Migration Squashing, Model Factory classes, Tailwind CSS (Used for Pagination Views), and usability improvements.</span></p>
<p style="margin-bottom: 0.35cm;line-height: 115%;text-align: left;background: transparent;"><br>&nbsp;</p>
<h3 style='margin-top: 0.35cm;margin-bottom: 0cm;color: #4f81bd;line-height: 115%;text-align: left;background: transparent;font-family: "Cambria", serif;font-weight: bold;'><strong><span lang="en-US"><u>5) How to install laravel via composer ?</u></span></strong><u>&nbsp;</u></h3>
<p style="margin-bottom: 0.35cm;line-height: 115%;text-align: left;background: transparent;"><br>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><span lang="en-US">The installation of Laravel via composer is very easy. You can install Laravel via composer by running the below command in the command prompt.</span></p>
<p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;"><span lang="en-US">composer create-project laravel/laravel your-project-name version</span></p>
<p style="margin-bottom: 0.35cm;line-height: 115%;text-align: left;background: transparent;"><br>&nbsp;</p>
<p style="margin-bottom: 0.35cm;line-height: 115%;text-align: left;background: transparent;"><br>&nbsp;</p>
<h3 style='margin-top: 0.35cm;margin-bottom: 0cm;color: #4f81bd;line-height: 115%;text-align: left;background: transparent;font-family: "Cambria", serif;font-weight: bold;'><strong><span lang="en-US"><u>6) List some features of laravel 6 ?</u></span></strong><span lang="en-US"><u>&nbsp;</u></span></h3>
<p style="margin-bottom: 0.35cm;line-height: 115%;text-align: left;background: transparent;"><br>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><span lang="en-US">Laravel 6 features</span></p>
<ul>
    <li>
        <p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">Inbuilt CRSF (cross-site request forgery ) Protection.</p>
    </li>
    <li>
        <p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;">Inbuilt paginations</p>
    </li>
    <li>
        <p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;"><u><span lang="en-US">Reverse Routing</span></u></p>
    </li>
    <li>
        <p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;">Query builder</p>
    </li>
    <li>
        <p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;">Route caching</p>
    </li>
    <li>
        <p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;">Database Migration</p>
    </li>
    <li>
        <p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;">IOC (Inverse of Control) Container Or service container.</p>
    </li>
    <li>
        <p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;">Job middleware</p>
    </li>
    <li>
        <p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;">Lazy collections</p>
    </li>
</ul>
<p style="margin-bottom: 0cm;line-height: 115%;text-align: left;background: transparent;"><br></p>
<h3 style='margin-top: 0.35cm;margin-bottom: 0cm;color: #4f81bd;line-height: 115%;text-align: left;background: transparent;font-family: "Cambria", serif;font-weight: bold;'><strong><span lang="en-US"><u>7) What is PHP artisan. List out some artisan commands ?</u></span></strong><span lang="en-US"><u>&nbsp;</u></span></h3>
<p style="margin-bottom: 0.35cm;line-height: 115%;text-align: left;background: transparent;"><br>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong><span lang="en-US">PHP artisan</span></strong><span lang="en-US">&nbsp;is the command line interface/tool included with Laravel. It provides a number of helpful commands that can help you while you build your application easily. Here are the list of some artisan command:-</span></p>
<ul>
    <li>
        <p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">php artisan list</p>
    </li>
    <li>
        <p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;">php artisan help</p>
    </li>
    <li>
        <p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;">php artisan tinker</p>
    </li>
    <li>
        <p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;">php artisan make</p>
    </li>
    <li>
        <p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;">php artisan &ndash;versian</p>
    </li>
    <li>
        <p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;">php artisan make model model_name</p>
    </li>
    <li>
        <p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;">php artisan make controller controller_name</p>
    </li>
</ul>
<p style="margin-bottom: 0cm;line-height: 115%;text-align: left;background: transparent;"><br></p>
<h3 style='margin-top: 0.35cm;margin-bottom: 0cm;color: #4f81bd;line-height: 115%;text-align: left;background: transparent;font-family: "Cambria", serif;font-weight: bold;'><strong><span lang="en-US"><u>8) List some default packages provided by Laravel Framework?</u></span></strong><span lang="en-US"><u>&nbsp;</u></span></h3>
<p style="margin-bottom: 0.35cm;line-height: 115%;text-align: left;background: transparent;"><br>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><span lang="en-US">Below are a list of some official/ default packages provided by Laravel&nbsp;</span></p>
<ul>
    <li>
        <p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">Cashier</p>
    </li>
    <li>
        <p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;">Envoy</p>
    </li>
    <li>
        <p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;">Passport</p>
    </li>
    <li>
        <p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;">Scout</p>
    </li>
    <li>
        <p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;">Socialite</p>
    </li>
    <li>
        <p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;">Horizon</p>
    </li>
    <li>
        <p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;">Telescope</p>
    </li>
</ul>
<p style="margin-bottom: 0cm;line-height: 115%;text-align: left;background: transparent;"><br></p>
<h3 style='margin-top: 0.35cm;margin-bottom: 0cm;color: #4f81bd;line-height: 115%;text-align: left;background: transparent;font-family: "Cambria", serif;font-weight: bold;'><strong><span lang="en-US"><u>9) What are named routes in Laravel?</u></span></strong><span lang="en-US"><u>&nbsp;</u></span></h3>
<p style="margin-bottom: 0.35cm;line-height: 115%;text-align: left;background: transparent;"><br>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong><span lang="en-US">Named routing</span></strong><span lang="en-US">&nbsp;is another amazing feature of Laravel framework. Named routes allow referring to routes when generating redirects or Urls more comfortably.<br>&nbsp;</span><strong><span lang="en-US">You can specify named routes by chaining the name method onto the route definition:</span></strong></p>
<p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;"><span lang="en-US">Route::get(&apos;user/profile&apos;, function () {</span></p>
<p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;"><span lang="en-US">//</span></p>
<p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;"><span lang="en-US">})-&gt;name(&apos;profile&apos;);</span></p>
<p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;"><br></p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong><span lang="en-US">You can specify route names for controller actions:</span></strong></p>
<p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;"><span lang="en-US">Route::get(&apos;user/profile&apos;, &apos;UserController@@showProfile&apos;)-&gt;name(&apos;profile&apos;);</span></p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><strong><span lang="en-US">Once you have assigned a name to your routes, you may use the route&apos;s name when generating URLs or redirects via the global route function:</span></strong></p>
<p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;"><span lang="en-US">// Generating URLs...</span></p>
<p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;"><span lang="en-US">$url = route(&apos;profile&apos;);</span></p>
<p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;"><span lang="en-US">// Generating Redirects...</span></p>
<p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;"><span lang="en-US">return redirect()-&gt;route(&apos;profile&apos;);</span></p>
<p style="margin-bottom: 0.35cm;line-height: 115%;text-align: left;background: transparent;"><br>&nbsp;</p>
<h3 style='margin-top: 0.35cm;margin-bottom: 0cm;color: #4f81bd;line-height: 115%;text-align: left;background: transparent;font-family: "Cambria", serif;font-weight: bold;'><strong><span lang="en-US"><u>10) What are the best features of Laravel 8?</u></span></strong><span lang="en-US"><u>&nbsp;</u></span></h3>
<p style="margin-bottom: 0.35cm;line-height: 115%;text-align: left;background: transparent;"><br>&nbsp;</p>
<p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;"><span lang="en-US">Laravel 8 was released with the new features. The best features of Laravel 8 are as follows -</span></p>
<ul>
    <li>
        <p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;margin-top: 0.49cm;">app/Models Directory</p>
    </li>
    <li>
        <p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;">New Landing Page</p>
    </li>
    <li>
        <p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;">Route Caching</p>
    </li>
    <li>
        <p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;">Maintenance Mode</p>
    </li>
    <li>
        <p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;">Job Batching</p>
    </li>
    <li>
        <p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;">Laravel Jetstream</p>
    </li>
    <li>
        <p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;">Controllers Routing Namespacing</p>
    </li>
    <li>
        <p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;">Better Syntax for Event Listening</p>
    </li>
    <li>
        <p style="margin-bottom: 0cm;line-height: 100%;text-align: left;background: transparent;">Queueable Anonymous Event Listeners</p>
    </li>
    <li>
        <p style="margin-bottom: 0.49cm;line-height: 100%;text-align: left;background: transparent;">Attributes on Extended Blade Components.</p>
    </li>
</ul>

</div>
</section>


@endsection