<?php
// database\seeders\userSeeder.php
namespace Database\Seeders;

use Illuminate\Database\Seeder;
use App\Models\user;  // to be added

class UserSeeder extends Seeder
{
    /**
     * Run the database seeds.
     *
     * @return void
     */
    public function run()
    {
      // Call for the model
      user::factory()
      ->count(4)
      ->create();
    }
}

