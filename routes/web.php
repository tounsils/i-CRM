<?php

use Illuminate\Support\Facades\Route;
use App\Http\Controllers\CustomAuthController;

/*
|--------------------------------------------------------------------------
| Web Routes
|--------------------------------------------------------------------------
|
| Here is where you can register web routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| contains the "web" middleware group. Now create something great!
| in laravel Route there is two parameters url and call back 
|
*/
/*
Route::get('/', function () {
    return view('welcome');
});
*/

Route::get('/', 'App\Http\Controllers\PresentationController@index');
Route::get('/about', 'App\Http\Controllers\PresentationController@about');
Route::get('/LaravelQuestions', 'App\Http\Controllers\PresentationController@LaravelQuestions');
Route::get('/LaravelDatabaseSeeding', 'App\Http\Controllers\PresentationController@LaravelDatabaseSeeding');
Route::get('/LaravelTinker', 'App\Http\Controllers\PresentationController@LaravelTinker');

Route::get('/blog', 'App\Http\Controllers\BlogController@index');

Route::get('/hello', function () {
    return "This is a simple Route returning hello message";
});

Route::get('users/{id}', function ($id) {
    return "Your id id $id";
});

/*
Route::get('/usershome', 'App\Http\Controllers\usersController@index')->name('usershome');
Route::get('/usercreate', 'App\Http\Controllers\usersController@create')->name('usercreate');
Route::post('/usercreate', 'App\Http\Controllers\usersController@store')->name('store');
*/
Route::resource('users', 'App\Http\Controllers\usersController');
Route::get('users', [App\Http\Controllers\usersController::class, 'getData']);

//Route::get('/companyhome', 'App\Http\Controllers\companyController@index')->name('companyhome');
//Route::resource('companyhome', 'App\Http\Controllers\companyController');
//Route::resource('customers','CustomerController');
//Route::resource('companyhome','companyController');
Route::resource('company', 'App\Http\Controllers\companyController');
Route::get('company', [App\Http\Controllers\companyController::class, 'getData']);

//Route::get('welcome', [CustomAuthController::class, 'welcome']); 
Route::get('login', [CustomAuthController::class, 'index'])->name('login');
Route::post('custom-login', [CustomAuthController::class, 'customLogin'])->name('login.custom'); 
Route::get('registration', [CustomAuthController::class, 'registration'])->name('register-user');
Route::post('custom-registration', [CustomAuthController::class, 'customRegistration'])->name('register.custom'); 
Route::get('signout', [CustomAuthController::class, 'signOut'])->name('signout');

Route::get('welcome', function () {
    return view('welcome');
});
