<?php
// database\seeders\DatabaseSeeder.php
namespace Database\Seeders;

use Illuminate\Database\Seeder;
use Illuminate\Database\Factories; 
use App\Models\appusers;

class DatabaseSeeder extends Seeder
{
    /**
     * Seed the application's database.
     *
     * @return void
     */
    public function run()
    {
      $this->call([
        
         AdminSeeder::class,
    ]);

      }
}
 //      CompanySeeder::class,
 //      UserSeeder::class,
 // RoleSeeder::class,
